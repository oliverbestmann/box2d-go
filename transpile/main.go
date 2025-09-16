package main

import (
	"bufio"
	"flag"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
	ccgo "modernc.org/ccgo/v4/lib"
)

// PrefixField is a unicode F symbol that is not part of the box2d code.
// We use this to easily identify & capitalize field names
const PrefixField = "ｆ"

var PREDEF_ATOMIC = `
#include <stdbool.h>
#include <stdint.h>
bool __builtin___atomic_compare_exchange_n(int32_t *ptr, int32_t *expected, int32_t desired, bool weak, int32_t success_memorder, int32_t failure_memorder);
`

var PREDEF_ALLOC = `
#include <stddef.h>
#define aligned_alloc __builtin_aligned_alloc
extern void* __builtin_aligned_alloc(size_t align, size_t size);
`

const PKGNAME = "b2"

func main() {
	// defer profile.Start(profile.CPUProfile).Stop()

	var noTranspile bool

	flag.BoolVar(&noTranspile, "no-transpile", false, "do not transpile the c sources")
	flag.Parse()

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	files, err := filepath.Glob("../box2d-c/src/*.c")
	must(err)

	args := []string{
		"ccgo",
		"-verify-types",
		"--package-name", PKGNAME,
		"-D", "BOX2D_DISABLE_SIMD",
		// "-D", "NDEBUG",
		"--predef", PREDEF_ATOMIC,
		"--predef", PREDEF_ALLOC,
		"--prefix-field", PrefixField,
		"-o", "../box2d_c.go",
		"-I", "../box2d-c/include/",
	}

	args = append(args, files...)

	if !noTranspile {
		// remove previous files
		gofs, err := filepath.Glob("../box2d_c*.go")
		must(err)
		for _, gof := range gofs {
			must(os.Remove(gof))
		}

		slog.Info("Transpile c to go")
		must(ccgo.NewTask(goos, goarch, args, os.Stdout, os.Stderr, nil).Main())
	}

	slog.Info("Read public API from header files")
	api := readAPI()

	slog.Info("Read generated go file")
	bytesSrc, err := os.ReadFile("../box2d_c.go")
	must(err)

	source := string(bytesSrc)

	slog.Info("Perform code cleanup via regexp")
	source = regexp.MustCompile("//go:build .*").ReplaceAllLiteralString(source, "")

	source = regexp.MustCompile(PrefixField+"([A-Za-z])([A-Za-z0-9_]*)").ReplaceAllStringFunc(source, func(s string) string {
		name := strings.TrimPrefix(s, PrefixField)
		name = strings.ToUpper(name[:1]) + name[1:]
		return name
	})

	slog.Info("Parse generated go source")
	fset := token.NewFileSet()
	var root dst.Node
	root, err = decorator.ParseFile(fset, "box2d_c.go", source, parser.ParseComments|parser.SkipObjectResolution)
	must(err)

	slog.Info("Apply code migrations")
	dst.Inspect(root, renameTLSType())
	root = dstutil.Apply(root, nil, removeModernC())
	root = dstutil.Apply(root, nil, hideSupportFuncs())
	root = dstutil.Apply(root, nil, fixQSortInvocation())
	root = dstutil.Apply(root, nil, useApiTypes(api.ExposedTypes))
	root = dstutil.Apply(root, nil, forwardDefaultAssertFcnToGo())
	dst.Inspect(root, renameApiTypes(api.ExposedTypes))
	dst.Inspect(root, removeTypeAliases())
	hideConsts(root)

	// dst.Inspect(root, injectCheckptr())

	var decls []dst.Decl

	createClassTypes(&decls)

	slog.Info("Generate go source for box2d api")
	dst.Inspect(root, createWrappers(&decls, api))
	dst.Inspect(root, exportVars(&decls, api.Exposed))

	writeSource("../box2d_api.go", &dst.File{
		Name:  dst.NewIdent(PKGNAME),
		Decls: withImport(decls, "unsafe", "runtime"),
	})

	for idx := range extracted {
		root = dstutil.Apply(root, nil, extractInit(&extracted[idx].Decls, extracted[idx].Pattern))
	}

	writeSource("../box2d_c.go", root.(*dst.File))

	for _, ex := range extracted {
		writeSource("../"+ex.Name, &dst.File{
			Name:  dst.NewIdent(PKGNAME),
			Decls: withImport(ex.Decls, "unsafe", "reflect"),
		})
	}
}

func hideConsts(root dst.Node) {
	allConsts := map[string]bool{}

	dst.Inspect(root, func(node dst.Node) bool {
		if decl, ok := node.(*dst.GenDecl); ok {
			return decl.Tok == token.CONST
		}

		if n, ok := node.(*dst.ValueSpec); ok {
			for _, name := range n.Names {
				if name.IsExported() {
					allConsts[name.Name] = true
				}
			}
		}

		return true
	})

	dst.Inspect(root, func(node dst.Node) bool {
		if ident, ok := node.(*dst.Ident); ok {
			if allConsts[ident.Name] {
				ident.Name = "_" + ident.Name
			}
		}

		return true
	})
}

func injectCheckptr() Visitor {
	var injectCheckptrInner Visitor = func(node dst.Node) bool {
		// unsafe.Pointer(...)
		ptrCall, ok := node.(*dst.CallExpr)
		if !ok {
			return true
		}

		// unsafe.Pointer
		fnName, ok := ptrCall.Fun.(*dst.SelectorExpr)
		if !ok {
			return true
		}

		if !(identOf(fnName.X) == "unsafe" && fnName.Sel.Name == "Pointer") {
			return true
		}

		// found the pointer, wrap with call to checkAddr
		wrapArg := &dst.CallExpr{
			Fun:  dst.NewIdent("checkAddr"),
			Args: []dst.Expr{dst.NewIdent("tls"), ptrCall.Args[0]},
		}

		ptrCall.Args[0] = wrapArg

		return true
	}

	return func(node dst.Node) bool {
		fn, ok := node.(*dst.FuncDecl)
		if ok && fn.Type.Params != nil && len(fn.Type.Params.List) > 0 {
			if star, ok := fn.Type.Params.List[0].Type.(*dst.StarExpr); ok {
				if identOf(star.X) == "_Stack" {
					dst.Inspect(node, injectCheckptrInner)
				}

				return false
			}
		}

		return true
	}
}

type extract struct {
	Pattern *regexp.Regexp
	Name    string
	Decls   []dst.Decl
}

var extracted = []extract{
	{
		Pattern: regexp.MustCompile("^init$"),
		Name:    "box2d_c-init.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*BroadPhase"),
		Name:    "box2d_c-broad.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*World"),
		Name:    "box2d_c-world.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*Body"),
		Name:    "box2d_c-body.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*Shape"),
		Name:    "box2d_c-shape.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*Joint"),
		Name:    "box2d_c-joint.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*Solve"),
		Name:    "box2d_c-solve.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*Sensor"),
		Name:    "box2d_c-sensor.go",
	},
	{
		Pattern: regexp.MustCompile("^b2Collide"),
		Name:    "box2d_c-collide.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*DynamicTree"),
		Name:    "box2d_c-tree.go",
	},
	{
		Pattern: regexp.MustCompile("^b2[^_]*Compute"),
		Name:    "box2d_c-compute.go",
	},
}

func withImport(decls []dst.Decl, imports ...string) []dst.Decl {
	var prelude []dst.Decl

	d := &dst.GenDecl{Tok: token.IMPORT}
	prelude = append(prelude, d)

	for _, i := range imports {
		d.Specs = append(d.Specs, &dst.ImportSpec{
			Path: &dst.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote(i),
			},
		})
	}

	useImport := func(pkg, ty string) {
		prelude = append(prelude, &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names: []*dst.Ident{dst.NewIdent("_")},
					Type:  &dst.SelectorExpr{X: dst.NewIdent(pkg), Sel: dst.NewIdent(ty)},
				},
			},
		})
	}

	for _, i := range imports {
		switch i {
		case "unsafe":
			useImport("unsafe", "Pointer")

		case "reflect":
			useImport("reflect", "Type")

		case "runtime":
			useImport("runtime", "MemStats")
		}
	}

	return append(prelude, decls...)
}

func extractInit(extracted *[]dst.Decl, pattern *regexp.Regexp) dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if fn, ok := cursor.Node().(*dst.FuncDecl); ok {
			if !pattern.MatchString(fn.Name.Name) {
				return true
			}

			cursor.Delete()

			*extracted = append(*extracted, fn)
		}

		return true
	}
}

func writeSource(target string, src *dst.File) {
	slog.Info("Write " + target)

	fp, err := os.Create(target)
	must(err)
	defer fp.Close()

	w := bufio.NewWriter(fp)
	defer func() { must(w.Flush()) }()

	must(decorator.Fprint(w, src))
}

func useApiTypes(api map[string]bool) dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if field, ok := cursor.Node().(*dst.Field); ok {
			field.Type = renameApiIdent(field, api)
		}

		return true
	}
}

func removeModernC() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		n, ok := cursor.Node().(*dst.ImportSpec)
		if !ok {
			return true
		}

		if strings.Contains(n.Path.Value, "modernc.org/") {
			cursor.Delete()
		}

		return true
	}
}

type Visitor = func(dst.Node) bool

func renameApiTypes(api map[string]bool) Visitor {
	return func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.Ident:
			if api[node.Name] {
				node.Name = node.Name[2:]
			}
		}

		return true
	}
}

func removeTypeAliases() Visitor {
	return func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.TypeSpec:
			_, ok := node.Type.(*dst.StructType)
			if ok {
				node.Assign = false
			}
		}

		return true
	}
}

func renameTLSType() Visitor {
	return func(node dst.Node) bool {
		switch node := node.(type) {
		case *dst.Ident:
			if node.Name == "TLS" {
				node.Name = "_Stack"
			}
		}

		return true
	}
}

func exportVars(decls *[]dst.Decl, api map[string]bool) Visitor {
	return func(node dst.Node) bool {
		if co, ok := node.(*dst.GenDecl); ok && (co.Tok == token.CONST || co.Tok == token.VAR) {
			for _, spec := range co.Specs {
				if dec, ok := spec.(*dst.ValueSpec); ok {
					ident := dec.Names[0]

					if api[ident.Name] {
						if !strings.HasPrefix(ident.Name, "b2_") {
							panic("expected ident to have b2_ prefix, got " + ident.Name)
						}

						name := ident.Name[3:]
						name = strings.ToUpper(name[:1]) + name[1:]

						if strings.HasSuffix(name, "Count") {
							continue
						}

						var ty dst.Expr

						switch {
						case strings.HasPrefix(name, "Color"):
							ty = dst.NewIdent("HexColor")

						case strings.HasSuffix(name, "Joint"):
							ty = dst.NewIdent("JointType")
							name = "JointType" + strings.TrimSuffix(name, "Joint")

						case strings.HasSuffix(name, "Body"):
							ty = dst.NewIdent("BodyType")

						case strings.HasSuffix(name, "Shape"):
							ty = dst.NewIdent("ShapeType")

						case strings.HasPrefix(name, "ToiState"):
							ty = dst.NewIdent("TOIState")

						default:
							panic("unknown enum type for " + name)
						}

						*decls = append(*decls, &dst.GenDecl{
							Tok: co.Tok,
							Specs: []dst.Spec{
								&dst.ValueSpec{
									Names:  []*dst.Ident{dst.NewIdent(name)},
									Values: []dst.Expr{dst.NewIdent(ident.Name)},
									Type:   ty,
								},
							},
						})
					}
				}
			}
		}

		return true
	}
}

func createClassTypes(decls *[]dst.Decl) {
	for _, ty := range classTypes {
		var fields dst.FieldList

		if ty.Extends == "" {
			fields.List = append(fields.List, &dst.Field{
				Names: []*dst.Ident{dst.NewIdent("Id")},
				Type:  dst.NewIdent(ty.IdType),
			})

			// fields.List = append(fields.List, &dst.Field{
			// 	Names: []*dst.Ident{dst.NewIdent("tls")},
			// 	Type:  &dst.StarExpr{X: dst.NewIdent("_Stack")},
			// })
		} else {
			fields.List = append(fields.List, &dst.Field{
				Type: dst.NewIdent(ty.Extends),
			})
		}

		*decls = append(*decls, &dst.GenDecl{
			Tok: token.TYPE,
			Specs: []dst.Spec{
				&dst.TypeSpec{
					Name: dst.NewIdent(ty.Name),
					Type: &dst.StructType{
						Fields: &fields,
					},
				},
			},
		})
	}
}

func createWrappers(decls *[]dst.Decl, api API) Visitor {
	return func(node dst.Node) bool {
		decl, ok := node.(*dst.FuncDecl)
		if !ok {
			return true
		}

		if !api.Exposed[decl.Name.Name] {
			return true
		}

		var classType *ClassTypes

		for idx := range classTypes {
			ty := &classTypes[idx]
			//if strings.HasPrefix(decl.Name.Name, ty.FnPrefix) {
			if decl.Type.Params != nil && len(decl.Type.Params.List) >= 2 {
				arg0, ok := decl.Type.Params.List[1].Type.(*dst.Ident)
				if ok && arg0.Name == ty.IdType {
					classType = ty
					break
				}
			}
			//}
		}

		var args []dst.Expr

		args = append(args, dst.NewIdent("theStack"))

		var body []dst.Stmt

		var paramTypes []*dst.Field
		for _, param := range decl.Type.Params.List[1:] {
			for _, name := range param.Names {
				var arg dst.Expr = dst.NewIdent(name.Name)
				name = dst.NewIdent(name.Name)

				paramType := renameApiIdent(param, api.ExposedTypes)

				if isUintptr(param.Type) && api.PointerTypes[decl.Name.Name][name.Name].IsValid() {
					actualType := api.PointerTypes[decl.Name.Name][name.Name]

					if actualType.Const {
						paramType = dst.NewIdent(actualType.GoType(&api))

						stackName := name.Name + "Stack"
						body = append(body,
							&dst.AssignStmt{
								Tok: token.DEFINE,
								Lhs: []dst.Expr{dst.NewIdent(stackName)},
								Rhs: []dst.Expr{
									&dst.CallExpr{
										Fun: dst.NewIdent("copyToStack"),
										Args: []dst.Expr{
											dst.NewIdent("theStack"),
											dst.NewIdent(name.Name),
										},
									},
								},
							},
							&dst.DeferStmt{
								Call: &dst.CallExpr{
									Fun: &dst.SelectorExpr{
										X:   dst.NewIdent(stackName),
										Sel: dst.NewIdent("Free"),
									},
								},
							},
						)

						arg = &dst.SelectorExpr{
							X:   dst.NewIdent(stackName),
							Sel: dst.NewIdent("Addr"),
						}
					} else {
						paramType = &dst.StarExpr{
							X: dst.NewIdent(actualType.GoType(&api)),
						}

						body = append(body, &dst.ExprStmt{
							X: &dst.CallExpr{
								Fun:  dst.NewIdent("escapes"),
								Args: []dst.Expr{dst.NewIdent(name.Name)},
							},
						})

						// prepend defer
						body = append([]dst.Stmt{
							&dst.DeferStmt{
								Call: &dst.CallExpr{
									Fun: &dst.SelectorExpr{
										X:   dst.NewIdent("runtime"),
										Sel: dst.NewIdent("KeepAlive"),
									},
									Args: []dst.Expr{dst.NewIdent(name.Name)},
								},
							},
						}, body...)

						arg = &dst.CallExpr{
							Fun: dst.NewIdent("uintptr"),
							Args: []dst.Expr{
								&dst.CallExpr{
									Fun: &dst.SelectorExpr{
										X:   dst.NewIdent("unsafe"),
										Sel: dst.NewIdent("Pointer"),
									},

									Args: []dst.Expr{dst.NewIdent(name.Name)},
								},
							},
						}
					}
				}

				if classType := getClassTypeById(identOf(paramType)); classType != nil {
					name.Name = strings.TrimSuffix(name.Name, "Id")
					paramType = dst.NewIdent(classType.Name)
					arg = &dst.SelectorExpr{
						X:   arg,
						Sel: dst.NewIdent("Id"),
					}
				}

				paramTypes = append(paramTypes, &dst.Field{
					Names: []*dst.Ident{dst.NewIdent(name.Name)},
					Type:  paramType,
				})

				args = append(args, arg)
			}
		}

		var returnTypes []*dst.Field
		if decl.Type.Results != nil {
			for _, param := range decl.Type.Results.List {
				ty := renameApiIdent(param, api.ExposedTypes)

				wrapTy := getClassTypeById(identOf(ty))
				if wrapTy != nil {
					ty = dst.NewIdent(wrapTy.Name)
				}

				if identOf(ty) == "Joint" && strings.HasPrefix(decl.Name.Name, "b2Create") {
					jointType := strings.TrimPrefix(decl.Name.Name, "b2Create")
					ty = dst.NewIdent(jointType)
				}

				for _, name := range param.Names {
					returnTypes = append(returnTypes, &dst.Field{
						Type:  ty,
						Names: []*dst.Ident{dst.NewIdent(name.Name)},
					})
				}
			}
		}

		call := &dst.CallExpr{
			Fun:  decl.Name,
			Args: args,
		}

		if decl.Type.Results != nil && len(decl.Type.Results.List) > 0 {
			assign := &dst.AssignStmt{
				Tok: token.ASSIGN,
				Rhs: []dst.Expr{call},
			}

			for _, ret := range decl.Type.Results.List {
				for _, name := range ret.Names {
					isWrapper := getClassTypeById(identOf(ret.Type)) != nil
					if isWrapper {
						assign.Lhs = append(assign.Lhs, &dst.SelectorExpr{
							X:   dst.NewIdent(name.Name),
							Sel: dst.NewIdent("Id"),
						})

						//	// body.tls = b.tls
						//	body = append(body, &dst.AssignStmt{
						//		Tok: token.ASSIGN,
						//		Lhs: []dst.Expr{
						//			&dst.SelectorExpr{
						//				X:   dst.NewIdent(name.Name),
						//				Sel: dst.NewIdent("tls"),
						//			},
						//		},
						//		Rhs: []dst.Expr{
						//			&dst.SelectorExpr{
						//				X:   dst.NewIdent("b"),
						//				Sel: dst.NewIdent("tls"),
						//			},
						//		},
						//	})

					} else {
						assign.Lhs = append(assign.Lhs, dst.NewIdent(name.Name))
					}
				}
			}

			body = append(body, assign)

			body = append(body, &dst.ReturnStmt{})
		} else {
			body = append(body, &dst.ExprStmt{X: call})
		}

		var fdecl *dst.FuncDecl

		if classType != nil {
			// adjust the id parameter value to come from the wrapper type
			args[1] = &dst.SelectorExpr{
				X:   dst.NewIdent("b"),
				Sel: dst.NewIdent("Id"),
			}

			name := strings.TrimPrefix(decl.Name.Name, classType.FnPrefix)
			name = strings.TrimPrefix(name, "b2")

			fdecl = &dst.FuncDecl{
				Name: dst.NewIdent(name),
				Recv: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("b")},
							Type:  dst.NewIdent(classType.Name),
						},
					},
				},
				Type: &dst.FuncType{
					TypeParams: decl.Type.TypeParams,
					Params:     &dst.FieldList{List: paramTypes[1:]},
					Results:    &dst.FieldList{List: returnTypes},
				},
				Body: &dst.BlockStmt{List: body},
			}
		} else {
			slog.Info("Wrap normal method", slog.String("method", decl.Name.Name))

			// normal method
			fdecl = &dst.FuncDecl{
				Name: dst.NewIdent(strings.ReplaceAll(decl.Name.Name[2:], "_", "")),
				Type: &dst.FuncType{
					TypeParams: decl.Type.TypeParams,
					Params:     &dst.FieldList{List: paramTypes},
					Results:    &dst.FieldList{List: returnTypes},
				},
				Body: &dst.BlockStmt{List: body},
			}
		}

		*decls = append(*decls, fdecl)

		return true
	}
}

func needClassTypeWrap(list []*dst.Field) bool {
	for _, f := range list {
		if getClassTypeById(identOf(f.Type)) != nil {
			return true
		}
	}

	return false
}

func identOf(expr dst.Expr) string {
	if ident, ok := expr.(*dst.Ident); ok {
		return ident.Name
	}

	return ""
}

func isUintptr(expr dst.Expr) bool {
	ident, ok := expr.(*dst.Ident)
	return ok && ident.Name == "uintptr"
}

func hideSupportFuncs() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if node, ok := cursor.Node().(*dst.SelectorExpr); ok {
			if ident, ok := node.X.(*dst.Ident); ok {
				if ident.Name == "libc" || ident.Name == "iqlibc" {
					if node.Sel.Name[0] == 'X' {
						node.Sel.Name = node.Sel.Name[1:]
					}

					if node.Sel.Name != "tls" {
						node.Sel.Name = strings.ToLower(node.Sel.Name[:1]) + node.Sel.Name[1:]
					}

					cursor.Replace(node.Sel)
				}
			}
		}

		return true
	}
}

func fixQSortInvocation() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		if call, ok := cursor.Node().(*dst.CallExpr); ok {
			if name, ok := call.Fun.(*dst.Ident); ok {
				if name.Name == "qsort" {
					ccgoCall, ok := call.Args[4].(*dst.CallExpr)
					if ok {

						if ccgoCall.Fun.(*dst.Ident).Name != "__ccgo_fp" {
							panic("expected call to __ccgo_fp")
						}

						call.Args[4] = ccgoCall.Args[0]
					}
				}
			}
		}

		return true
	}
}

func renameApiIdent(param *dst.Field, api map[string]bool) dst.Expr {
	ty, ok := param.Type.(*dst.Ident)
	if !ok {
		return param.Type
	}

	name := cleanApiName(api, ty.Name)
	if api[name] {
		tyCopy := *ty
		tyCopy.Name = name[2:]
		ty = &tyCopy
	}

	if name, ok := renameSimpleType(ty.Name); ok {
		tyCopy := *ty
		tyCopy.Name = name
		ty = &tyCopy
	}

	return ty
}

type ParamType struct {
	CType string
	Const bool
}

func (p ParamType) GoType(api *API) string {
	if api.ExposedTypes[p.CType] {
		return p.CType[2:]
	}

	return p.CType
}

func (p ParamType) IsValid() bool {
	return strings.HasPrefix(p.CType, "b2")
}

type API struct {
	Exposed      map[string]bool
	ExposedTypes map[string]bool

	// function to param name to type
	PointerTypes map[string]map[string]ParamType
}

func (a *API) addPointerType(fn, param string, ty ParamType) {
	fnt := a.PointerTypes[fn]

	if fnt == nil {
		fnt = map[string]ParamType{}
		a.PointerTypes[fn] = fnt
	}

	fnt[param] = ty
}

func readAPI() API {
	var api API

	api.Exposed = map[string]bool{}
	api.ExposedTypes = map[string]bool{}
	api.PointerTypes = map[string]map[string]ParamType{}

	headers, err := filepath.Glob("../box2d-c/include/box2d/*.h")
	if err != nil {
		panic(err)
	}

	reType := regexp.MustCompile(`typedef (?:struct|enum) (b2[A-Za-z0-9_]+)`)
	reFunc := regexp.MustCompile(`B2_API \S+ (b2[A-Za-z0-9_]+)`)
	reEnumValues := regexp.MustCompile(`(?m)^\s*(b2_[A-Za-z0-9_]+)`)
	rePointerTypes := regexp.MustCompile(`(const )?([A-Za-z0-9]+)\* ([A-Za-z0-9]+)`)

	for _, header := range headers {
		src, err := os.ReadFile(header)
		must(err)

		lines := strings.Split(string(src), "\n")

		for _, match := range reType.FindAllStringSubmatch(string(src), -1) {
			api.ExposedTypes[match[1]] = true
			api.Exposed[match[1]] = true
		}

		for _, match := range reFunc.FindAllStringSubmatch(string(src), -1) {
			api.Exposed[match[1]] = true
		}

		for _, match := range reEnumValues.FindAllStringSubmatch(string(src), -1) {
			api.Exposed[match[1]] = true
		}

		for _, line := range lines {
			match := reFunc.FindStringSubmatch(line)
			if len(match) == 0 {
				continue
			}

			fn := match[1]

			for _, match := range rePointerTypes.FindAllStringSubmatch(line, -1) {
				co, ty, param := match[1], match[2], match[3]
				api.addPointerType(fn, param, ParamType{
					CType: ty,
					Const: co != "",
				})
			}
		}
	}

	api.Exposed["b2InternalAssertFcn"] = false

	api.Exposed["b2SetAssertFcn"] = false
	api.Exposed["b2SetAllocator"] = false

	api.Exposed["b2Body_SetName"] = false
	api.Exposed["b2Body_GetName"] = false

	api.Exposed["b2Body_GetJoints"] = false
	api.Exposed["b2Body_GetShapes"] = false
	api.Exposed["b2Chain_GetSegments"] = false
	api.Exposed["b2Shape_GetSensorOverlaps"] = false

	api.Exposed["b2World_CollideMover"] = false
	api.Exposed["b2World_SetRestitutionCallback"] = false
	api.Exposed["b2World_SetFrictionCallback"] = false
	api.Exposed["b2World_OverlapAABB"] = false
	api.Exposed["b2World_OverlapShape"] = false
	api.Exposed["b2World_SetCustomFilterCallback"] = false
	api.Exposed["b2World_SetPreSolveCallback"] = false
	api.Exposed["b2World_DumpMemoryStats"] = false

	api.ExposedTypes["b2DebugDraw"] = false
	api.Exposed["b2World_Draw"] = false

	api.ExposedTypes["b2BodyEvents"] = false
	api.Exposed["b2World_GetBodyEvents"] = false

	api.ExposedTypes["b2SensorEvents"] = false
	api.Exposed["b2World_GetSensorEvents"] = false

	api.ExposedTypes["b2ContactEvents"] = false
	api.Exposed["b2World_GetContactEvents"] = false

	api.Exposed["b2Joint_GetConstraintTuning"] = false

	// we do not really need those, do we?
	api.Exposed["b2GetMillisecondsAndReset"] = false
	api.Exposed["b2Hash"] = false
	api.Exposed["b2ComputeHull"] = false

	for key := range api.Exposed {
		if strings.Contains(key, "DynamicTree") {
			api.Exposed[key] = false
		}
	}

	return api
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var reSimpleType = regexp.MustCompile("u?int(?:8|16|32|64)_t")

func renameSimpleType(name string) (string, bool) {
	if reSimpleType.MatchString(name) {
		return strings.TrimSuffix(name, "_t"), true
	}

	return name, false
}

func forwardDefaultAssertFcnToGo() dstutil.ApplyFunc {
	return func(cursor *dstutil.Cursor) bool {
		fn, ok := cursor.Node().(*dst.FuncDecl)
		if !ok {
			return true
		}

		if fn.Name.Name == "b2DefaultAssertFcn" {
			fn.Body.List = []dst.Stmt{
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CallExpr{
							Fun: dst.NewIdent("b2DefaultAssertFcnGo"),
							Args: []dst.Expr{
								dst.NewIdent("tls"),
								dst.NewIdent("condition"),
								dst.NewIdent("fileName"),
								dst.NewIdent("lineNumber"),
							},
						},
					},
				},
			}
		}

		return true
	}
}

func cleanApiName(api map[string]bool, name string) string {
	clean := strings.TrimRight(name, "0123456789")

	if !api[name] && api[clean] {
		return clean
	}

	return name
}

type ClassTypes struct {
	Name     string
	IdType   string
	FnPrefix string
	Extends  string
}

var classTypes = []ClassTypes{
	{
		Name:     "Body",
		IdType:   "BodyId",
		FnPrefix: "b2Body_",
	},
	{
		Name:     "Shape",
		IdType:   "ShapeId",
		FnPrefix: "b2Shape_",
	},
	{
		Name:     "World",
		IdType:   "WorldId",
		FnPrefix: "b2World_",
	},
	{
		Name:     "Joint",
		IdType:   "JointId",
		FnPrefix: "b2Joint_",
	},
	{
		Name:     "Chain",
		IdType:   "ChainId",
		FnPrefix: "b2Chain_",
	},
	{
		Name:     "PrismaticJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2PrismaticJoint_",
	},
	{
		Name:     "MotorJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2MotorJoint_",
	},
	{
		Name:     "RevoluteJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2RevoluteJoint_",
	},
	{
		Name:     "DistanceJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2DistanceJoint_",
	},
	{
		Name:     "WheelJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2WheelJoint_",
	},
	{
		Name:     "WeldJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2WeldJoint_",
	},
	{
		Name:     "MouseJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2MouseJoint_",
	},
	{
		Name:     "FilterJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2FilterJoint_",
	},
}

func getClassTypeById(idType string) *ClassTypes {
	for idx := range classTypes {
		if classTypes[idx].IdType == idType {
			return &classTypes[idx]
		}
	}

	return nil
}
