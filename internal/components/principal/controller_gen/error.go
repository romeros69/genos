package controller_gen

import (
	"go/ast"
	"go/token"
)

func (cg *ControllerHTTPGenerator) genErrorResponseAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("http"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/gin-gonic/gin\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("errResponse"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											ast.NewIdent("Error"),
										},
										Type: ast.NewIdent("string"),
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("errorResponse"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							0: {
								Names: []*ast.Ident{
									ast.NewIdent("c"),
								},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("gin"),
										Sel: ast.NewIdent("Context"),
									},
								},
							},
							1: {
								Names: []*ast.Ident{
									ast.NewIdent("code"),
								},
								Type: ast.NewIdent("int"),
							},
							2: {
								Names: []*ast.Ident{
									ast.NewIdent("msg"),
								},
								Type: ast.NewIdent("string"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("c"),
									Sel: ast.NewIdent("AbortWithStatusJSON"),
								},
								Args: []ast.Expr{
									0: ast.NewIdent("code"),
									1: &ast.CompositeLit{
										Type: ast.NewIdent("errResponse"),
										Elts: []ast.Expr{
											ast.NewIdent("msg"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
