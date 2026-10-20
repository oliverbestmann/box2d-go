package main

import (
	"fmt"
	"go/token"

	"github.com/dave/dst"
)

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
