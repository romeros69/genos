package principal

import (
	"go/ast"
	"go/token"
	"strings"
)

type ListGenerator struct {
}

func (lg *ListGenerator) GenAST() map[string]*ast.File {
	return nil
}

// генерация базовго кода для файла контроллера (без самих ручек)
func (lg *ListGenerator) controllerBaseAST() *ast.File {
	moduleName := "lol"
	entityName := "Car"
	return &ast.File{
		Name: ast.NewIdent("http"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"net/http\"",
						},
					},
					2: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/internal/entity\"",
						},
					},
					3: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/internal/usecase\"",
						},
					},
					4: &ast.ImportSpec{
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
						Name: ast.NewIdent(entityName + "Routes"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											ast.NewIdent("uc"),
										},
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("usecase"),
											Sel: ast.NewIdent(string(strings.ToLower(entityName)[0])),
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("new" + entityName + "Routes"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							0: {
								Names: []*ast.Ident{
									ast.NewIdent("handler"),
								},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("gin"),
										Sel: ast.NewIdent("RouterGroup"),
									},
								},
							},
							1: {
								Names: []*ast.Ident{
									ast.NewIdent(string(strings.ToLower(entityName)[0])),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("usecase"),
									Sel: ast.NewIdent(entityName),
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent(string(strings.ToLower(entityName)[0]) + "r"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: ast.NewIdent(strings.ToLower(entityName) + "Routes"),
										Elts: []ast.Expr{
											&ast.KeyValueExpr{
												Key:   ast.NewIdent(string(strings.ToLower(entityName)[0])),
												Value: ast.NewIdent(string(strings.ToLower(entityName)[0])),
											},
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
