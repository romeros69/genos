package base

import (
	"go/ast"
	"go/token"
)

type RouterGenerator struct {
	moduleName     string
	fullPathToFile string
}

var _ BaseGenerator = (*RouterGenerator)(nil)

func NewRouterGenerator(moduleName string) *RouterGenerator {
	return &RouterGenerator{
		moduleName:     moduleName,
		fullPathToFile: "internal/controller/http/router.go",
	}
}

func (rg *RouterGenerator) FullPathToFile() string {
	return rg.fullPathToFile
}

func (rg *RouterGenerator) GenAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("http"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/gin-gonic/gin\"",
						},
					},
				},
			},
			1: &ast.FuncDecl{
				Name: ast.NewIdent("NewRouter"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("handler"),
								},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("gin"),
										Sel: ast.NewIdent("Engine"),
									},
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						0: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("_"),
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("handler"),
										Sel: ast.NewIdent("Group"),
									},
									Args: []ast.Expr{
										&ast.BasicLit{
											Kind:  token.STRING,
											Value: "\"/api/v1\"",
										},
									},
								},
							},
						},
						1: &ast.BlockStmt{
							Lbrace: 11,
							Rbrace: 13,
						},
					},
				},
			},
		},
	}
}
