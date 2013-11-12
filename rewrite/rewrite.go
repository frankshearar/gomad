package rewrite

import (
	"go/ast"
)

type MonadRewriter struct {
	TypeName string
}

func (v MonadRewriter) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.File:
		{
			n.Name.Name = n.Name.Name + "_" + v.TypeName
			valueIndex := -1
			for i, m := range n.Decls {
				d, ok := m.(*ast.GenDecl)
				if ok {
					for _, k := range d.Specs {
						t, ok := k.(*ast.TypeSpec)
						if ok {
							if t.Name.Name == "Value" {
								valueIndex = i
								break
							}
						}
					}
				}
			}
			if valueIndex >= 0 {
				n.Decls = append(n.Decls[:valueIndex], n.Decls[valueIndex+1:]...)
			}
		}
	case *ast.Field:
		{
			if f, ok := n.Type.(*ast.Ident); ok {
				if f.Name == "Value" {
					f.Name = v.TypeName
				}
			}
		}
	}
	return v
}
