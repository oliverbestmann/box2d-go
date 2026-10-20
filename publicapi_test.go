package b2

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// A uintptr in the public API lets users hand Go memory to box2d, which is easy to get wrong.
// The generator unexports or retypes such fields, see transpile/publicapi.go.
func TestPublicAPIHasNoUintptr(t *testing.T) {
	names, err := filepath.Glob("*.go")
	require.NoError(t, err)

	fset := token.NewFileSet()

	var found []string
	for _, name := range names {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}

		file, err := parser.ParseFile(fset, name, nil, parser.SkipObjectResolution)
		require.NoError(t, err)

		report := func(node ast.Node, what string) {
			found = append(found, fset.Position(node.Pos()).String()+": "+what)
		}

		for _, decl := range file.Decls {
			switch decl := decl.(type) {
			case *ast.FuncDecl:
				if decl.Name.IsExported() && exportedReceiver(decl.Recv) && hasUintptr(decl.Type) {
					report(decl, decl.Name.Name)
				}

			case *ast.GenDecl:
				for _, spec := range decl.Specs {
					switch spec := spec.(type) {
					case *ast.TypeSpec:
						if spec.Name.IsExported() && hasUintptr(spec.Type) {
							report(spec, spec.Name.Name)
						}

					case *ast.ValueSpec:
						for _, ident := range spec.Names {
							if ident.IsExported() && spec.Type != nil && hasUintptr(spec.Type) {
								report(spec, ident.Name)
							}
						}
					}
				}
			}
		}
	}

	assert.Empty(t, found)
}

func exportedReceiver(recv *ast.FieldList) bool {
	if recv == nil {
		return true
	}

	ty := recv.List[0].Type
	if star, ok := ty.(*ast.StarExpr); ok {
		ty = star.X
	}

	ident, ok := ty.(*ast.Ident)
	return ok && ident.IsExported()
}

// hasUintptr reports whether expr mentions uintptr, ignoring unexported struct fields.
func hasUintptr(expr ast.Node) bool {
	found := false

	ast.Inspect(expr, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.Ident:
			found = found || node.Name == "uintptr"

		case *ast.StructType:
			for _, field := range node.Fields.List {
				exported := len(field.Names) == 0
				for _, name := range field.Names {
					exported = exported || name.IsExported()
				}

				if exported && hasUintptr(field.Type) {
					found = true
				}
			}

			return false
		}

		return !found
	})

	return found
}
