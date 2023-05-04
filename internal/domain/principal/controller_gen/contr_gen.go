package controller_gen

import (
	"genos/internal/domain/dsl"
	"go/ast"
	"go/token"
	"strings"
)

type ControllerHTTPGenerator struct {
	entAST            *dsl.Ent
	moduleName        string
	fullPathToPackage string
}

func NewControllerHTTPGenerator(moduleName string) *ControllerHTTPGenerator {
	return &ControllerHTTPGenerator{
		moduleName:        moduleName,
		fullPathToPackage: "internal/controller/http",
	}
}

func (cg *ControllerHTTPGenerator) controllerBaseAST() *ast.File {
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
							Value: "\"" + cg.moduleName + "/internal/entity\"",
						},
					},
					3: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + cg.moduleName + "/internal/usecase\"",
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
						Name: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Routes"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
										},
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("usecase"),
											Sel: ast.NewIdent(cg.entAST.Name),
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("new" + cg.entAST.Name + "Routes"),
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
									ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("usecase"),
									Sel: ast.NewIdent(cg.entAST.Name),
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
							},
							Tok: token.AND,
							Rhs: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Routes"),
										Elts: []ast.Expr{
											&ast.KeyValueExpr{
												Key:   ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
												Value: ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
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

func (cg *ControllerHTTPGenerator) getFullPathToFile(nameEntity string) string {
	return strings.Join([]string{cg.fullPathToPackage, nameEntity, ".go"}, "")
}
