package controller_gen

import (
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genDTOControllerAST() *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Request"),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: func() []*ast.Field {
							res := make([]*ast.Field, 0)
							for i := 1; i < len(cg.entAST.Fields); i++ {
								res = append(res, &ast.Field{
									Names: []*ast.Ident{ast.NewIdent(cg.entAST.Fields[i].Name)},
									Type:  ast.NewIdent(util.TypesMap[cg.entAST.Fields[i].TokType]),
								})
							}
							return res
						}(),
					},
				},
			},
		},
	}
}
