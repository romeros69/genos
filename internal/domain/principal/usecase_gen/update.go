package usecase_gen

import (
	"go/ast"
	"strings"
)

func (ug *UseCaseGenerator) genUpdateUCAST() *ast.FuncDecl {
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
		Name: ast.NewIdent(strings.Join([]string{"Update", ug.entAST.Name}, "")),
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
							ast.NewIdent(strings.ToLower(ug.entAST.Name)),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(ug.entAST.Name),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
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
								Sel: ast.NewIdent(strings.Join([]string{"Update", ug.entAST.Name}, "")),
							},
							Args: []ast.Expr{
								0: ast.NewIdent("ctx"),
								1: ast.NewIdent(strings.ToLower(ug.entAST.Name)),
							},
						},
					},
				},
			},
		},
	}
}
