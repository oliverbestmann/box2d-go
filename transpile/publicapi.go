package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
)

// The public Go API must not contain a uintptr, it is too easy to pass Go memory to box2d that
// way. TestPublicAPIHasNoUintptr in the package root checks the result.

// hideInternals unexports the names that ccgo exports for C declarations without the b2
// prefix, like static helper functions and context structs.
func hideInternals(root *dst.File, api API) {
	hidden := map[string]bool{}
	hide := func(name string) {
		_, isFunc := api.Exposed["b2"+name]
		_, isType := api.ExposedTypes["b2"+name]
		if token.IsExported(name) && !isFunc && !isType {
			hidden[name] = true
		}
	}

	for _, decl := range root.Decls {
		switch decl := decl.(type) {
		case *dst.FuncDecl:
			if decl.Recv == nil {
				hide(decl.Name.Name)
			}

		case *dst.GenDecl:
			if decl.Tok == token.TYPE {
				for _, spec := range decl.Specs {
					hide(spec.(*dst.TypeSpec).Name.Name)
				}
			}
		}
	}

	// the renaming below only works by name, so a field must not have the same name
	dst.Inspect(root, func(node dst.Node) bool {
		var name string
		switch node := node.(type) {
		case *dst.SelectorExpr:
			name = node.Sel.Name
		case *dst.Field:
			if len(node.Names) > 0 {
				name = node.Names[0].Name
			}
		}

		if hidden[name] {
			panic(fmt.Sprintf("can not hide %s, a field has the same name", name))
		}

		return true
	})

	dst.Inspect(root, func(node dst.Node) bool {
		if ident, ok := node.(*dst.Ident); ok && hidden[ident.Name] {
			ident.Name = "_" + ident.Name
		}

		return true
	})
}

type fieldChange int

const (
	// unexport fields that the Go API sets, like the points in Body.CreateChain
	unexport fieldChange = iota + 1

	// toUint64 turns a void* that box2d never reads into an opaque number
	toUint64
)

// apiFields are the C pointer fields of public structs, by Go type and field name.
var apiFields = map[string]map[string]fieldChange{
	"BodyDef":       {"Name": unexport, "UserData": toUint64},
	"ShapeDef":      {"UserData": toUint64},
	"ChainDef":      {"Points": unexport, "PointCount": unexport, "Materials": unexport, "MaterialCount": unexport, "UserData": toUint64},
	"JointDef":      {"UserData": toUint64},
	"WorldDef":      {"FrictionCallback": unexport, "RestitutionCallback": unexport, "EnqueueTask": unexport, "FinishTask": unexport, "UserTaskContext": unexport, "UserData": toUint64},
	"BodyMoveEvent": {"UserData": toUint64},
	"JointEvent":    {"UserData": toUint64},
}

// fixApiFields applies apiFields. A field can have the same name in many structs, so this
// needs the types of the selectors.
func fixApiFields(root *dst.File) {
	restorer := decorator.NewRestorer()
	file, err := restorer.RestoreFile(root)
	must(err)

	// the hand-written files define _Stack and the support functions, without them the types of
	// many expressions are invalid. Errors from the not yet generated API are expected.
	files := []*ast.File{file}
	for _, name := range handWrittenFiles() {
		f, err := parser.ParseFile(restorer.Fset, name, nil, parser.SkipObjectResolution)
		must(err)
		files = append(files, f)
	}

	info := &types.Info{Defs: map[*ast.Ident]types.Object{}, Uses: map[*ast.Ident]types.Object{}}
	conf := types.Config{Importer: importer.Default(), Error: func(error) {}}
	pkg, _ := conf.Check(PKGNAME, restorer.Fset, files, info)

	fields := map[*types.Var]fieldChange{}
	for typeName, changes := range apiFields {
		obj := pkg.Scope().Lookup(typeName)
		if obj == nil {
			panic("type not found: " + typeName)
		}

		st := obj.Type().Underlying().(*types.Struct)
		for name, change := range changes {
			idx := 0
			for idx < st.NumFields() && st.Field(idx).Name() != name {
				idx++
			}

			if idx == st.NumFields() {
				panic(fmt.Sprintf("field not found: %s.%s", typeName, name))
			}

			fields[st.Field(idx)] = change
		}
	}

	dstNodes := map[ast.Node]dst.Node{}
	for d, a := range restorer.Ast.Nodes {
		dstNodes[a] = d
	}

	// the declarations and all uses of the fields
	changes := map[*dst.Ident]fieldChange{}
	for _, objects := range []map[*ast.Ident]types.Object{info.Defs, info.Uses} {
		for ident, obj := range objects {
			field, ok := obj.(*types.Var)
			if !ok || fields[field] == 0 {
				continue
			}

			changes[dstNodes[ident].(*dst.Ident)] = fields[field]
		}
	}

	isUint64 := func(expr dst.Expr) bool {
		sel, ok := expr.(*dst.SelectorExpr)
		return ok && changes[sel.Sel] == toUint64
	}

	toType := func(ty string, expr dst.Expr) dst.Expr {
		return &dst.CallExpr{Fun: dst.NewIdent(ty), Args: []dst.Expr{expr}}
	}

	// selectors that are written to, all other uses read the field
	written := map[*dst.SelectorExpr]bool{}

	dstutil.Apply(root, func(c *dstutil.Cursor) bool {
		switch node := c.Node().(type) {
		case *dst.AssignStmt:
			for idx, lhs := range node.Lhs {
				if !isUint64(lhs) {
					continue
				}

				if node.Tok != token.ASSIGN || len(node.Lhs) != len(node.Rhs) {
					panic("unexpected assignment to a uint64 field")
				}

				written[lhs.(*dst.SelectorExpr)] = true
				node.Rhs[idx] = toType("uint64", node.Rhs[idx])
			}

		case *dst.KeyValueExpr:
			if key, ok := node.Key.(*dst.Ident); ok && changes[key] == toUint64 {
				node.Value = toType("uint64", node.Value)
			}

		case *dst.UnaryExpr:
			if node.Op == token.AND && isUint64(node.X) {
				panic("address of a uint64 field")
			}

		case *dst.IncDecStmt:
			if isUint64(node.X) {
				panic("increment of a uint64 field")
			}
		}

		return true
	}, func(c *dstutil.Cursor) bool {
		switch node := c.Node().(type) {
		case *dst.SelectorExpr:
			if isUint64(node) && !written[node] {
				c.Replace(toType("uintptr", node))
			}

		case *dst.Field:
			if len(node.Names) > 0 && changes[node.Names[0]] == toUint64 {
				node.Type = dst.NewIdent("uint64")
			}
		}

		return true
	})

	for ident, change := range changes {
		if change == unexport {
			ident.Name = strings.ToLower(ident.Name[:1]) + ident.Name[1:]
		}
	}
}

// handWrittenFiles returns the non-test Go files of the package root that are not generated
// and match the current build.
func handWrittenFiles() []string {
	names, err := filepath.Glob("../*.go")
	must(err)

	var files []string
	for _, name := range names {
		base := filepath.Base(name)
		if strings.HasPrefix(base, "box2d_c") || base == "box2d_api.go" || strings.HasSuffix(base, "_test.go") {
			continue
		}

		if ok, err := build.Default.MatchFile("..", base); err != nil || !ok {
			continue
		}

		files = append(files, name)
	}

	return files
}
