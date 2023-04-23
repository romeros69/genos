package usecase_gen

import (
	"genos/internal/util"
	"go/ast"
	"strings"
)

func (ug *UseCaseGenerator) genReadUCAST() *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent(string(strings.ToLower(ug.entAST.Name)[0])),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(strings.Join([]string{ug.entAST.Name, "UseCase"}, "")),
					},
				},
			},
		},
		Name: ast.NewIdent(strings.Join([]string{"Get", ug.entAST.Name, "By", ug.entAST.Fields[0].Name}, "")),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Names: []*ast.Ident{
							ast.NewIdent("ctx"),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
					1: {
						Names: []*ast.Ident{
							ast.NewIdent(strings.ToLower(ug.entAST.Fields[0].Name)),
						},
						Type: ast.NewIdent(util.TypesMap[ug.entAST.Fields[0].TokType]),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(ug.entAST.Name),
						},
					},
					1: {
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X: &ast.SelectorExpr{
									X:   ast.NewIdent(string(strings.ToLower(ug.entAST.Name)[0])),
									Sel: ast.NewIdent("repo"),
								},
								Sel: ast.NewIdent(strings.Join([]string{"Get", ug.entAST.Name, "By", ug.entAST.Fields[0].Name}, "")),
							},
							Args: []ast.Expr{
								0: ast.NewIdent("ctx"),
								1: ast.NewIdent(strings.ToLower(ug.entAST.Fields[0].Name)),
							},
						},
					},
				},
			},
		},
	}
}
