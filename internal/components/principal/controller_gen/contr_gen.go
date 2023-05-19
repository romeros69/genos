package controller_gen

import (
	"genos/internal/components/dsl"
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
		fullPathToPackage: "internal/controller/http/",
	}
}

func (cg *ControllerHTTPGenerator) GetMapControllerAST(dslAST *dsl.AST, prevRouterAST *ast.File) map[string]*ast.File {
	resultMap := make(map[string]*ast.File, len(dslAST.Entities)+1)
	for _, entityAST := range dslAST.Entities {
		resultMap[cg.getFullPathToFile(strings.ToLower(entityAST.Name))] = cg.genControllerAST(&entityAST)
	}
	resultMap[cg.getFullPathToFile("error")] = cg.genErrorResponseAST()
	resultMap[cg.getFullPathToFile("router")] = cg.genCallHandlerRouterAST(dslAST, prevRouterAST)
	return resultMap
}

func (cg *ControllerHTTPGenerator) genControllerAST(entAST *dsl.Ent) *ast.File {
	cg.entAST = entAST
	resultAST := cg.controllerBaseAST()
	resultAST.Decls = append(resultAST.Decls, cg.genDTOControllerAST())
	for _, action := range entAST.Actions {
		funcDecl, _ := resultAST.Decls[2].(*ast.FuncDecl)
		funcDecl.Body.List = append(funcDecl.Body.List, cg.genRoutesForEndpoint(action))
		switch action {
		case "create":
			resultAST.Decls = append(resultAST.Decls, cg.genCreateControllerAST())
		case "read":
			resultAST.Decls = append(resultAST.Decls, cg.genReadControllerAST())
		case "update":
			resultAST.Decls = append(resultAST.Decls, cg.genUpdateControllerAST())
		case "delete":
			resultAST.Decls = append(resultAST.Decls, cg.genDeleteControllerAST())
		case "list":
			resultAST.Decls = append(resultAST.Decls, cg.genListControllerAST())
		}
	}
	return resultAST
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
							Value: "\"net/http\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + cg.moduleName + "/internal/entity\"",
						},
					},
					2: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + cg.moduleName + "/internal/usecase\"",
						},
					},
					3: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/gin-gonic/gin\"",
						},
					},
					4: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"strconv\"",
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
											Sel: ast.NewIdent(cg.entAST.Name + "Contract"),
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
									Sel: ast.NewIdent(cg.entAST.Name + "Contract"),
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
							Tok: token.DEFINE,
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

func (cg *ControllerHTTPGenerator) genRoutesForEndpoint(action string) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X: ast.NewIdent("handler"),
				Sel: ast.NewIdent(func() string {
					switch action {
					case "create":
						return "POST"
					case "read":
						return "GET"
					case "update":
						return "PUT"
					case "delete":
						return "DELETE"
					case "list":
						return "GET"
					}
					return ""
				}()),
			},
			Args: []ast.Expr{
				0: &ast.BasicLit{
					Kind: token.STRING,
					Value: func() string {
						if action == "list" || action == "create" {
							return "\"/" + strings.ToLower(cg.entAST.Name) + "\""
						}
						return "\"/" + strings.ToLower(cg.entAST.Name) + "/:" + strings.ToLower(cg.entAST.Fields[0].Name) + "\""
					}(),
				},
				1: &ast.SelectorExpr{
					X: ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
					Sel: ast.NewIdent(func() string {
						switch action {
						case "create":
							return "create" + cg.entAST.Name
						case "read":
							return "get" + cg.entAST.Name + "By" + cg.entAST.Fields[0].Name
						case "update":
							return "update" + cg.entAST.Name
						case "delete":
							return "delete" + cg.entAST.Name
						case "list":
							return "get" + cg.entAST.Name + "List"
						}
						return ""
					}()),
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) getFullPathToFile(nameEntity string) string {
	return strings.Join([]string{cg.fullPathToPackage, nameEntity, ".go"}, "")
}
