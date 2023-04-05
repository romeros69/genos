package model

import (
	"go/ast"
	"go/token"
)

type AppGenerator struct {
	fullPathToFile string
	moduleName     string
}

var _ BaseGenerator = (*AppGenerator)(nil)

func NewAppGenerator(moduleName string) *AppGenerator {
	return &AppGenerator{
		moduleName:     moduleName,
		fullPathToFile: "internal/app/app.go",
	}
}

func (ag *AppGenerator) FullPathToFile() string {
	return ag.fullPathToFile
}

func (ag *AppGenerator) GenAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("app"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"log\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"os\"",
						},
					},
					2: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"os/signal\"",
						},
					},
					3: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + ag.moduleName + "/configs\"",
						},
					},
					4: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + ag.moduleName + "/pkg/httpserver\"",
						},
					},
					5: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + ag.moduleName + "/pkg/postgres\"",
						},
					},
					6: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"syscall\"",
						},
					},
					7: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"time\"",
						},
					},
					8: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/gin-contrib/cors\"",
						},
					},
					9: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/gin-gonic/gin\"",
						},
					},
				},
			},
			1: &ast.FuncDecl{
				Name: ast.NewIdent("Run"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("cfg"),
								},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("configs"),
										Sel: ast.NewIdent("Config"),
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
								0: ast.NewIdent("_"),
								1: ast.NewIdent("err"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("postgres"),
										Sel: ast.NewIdent("New"),
									},
									Args: []ast.Expr{
										ast.NewIdent("cfg"),
									},
								},
							},
						},
						1: &ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X:  ast.NewIdent("err"),
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("log"),
												Sel: ast.NewIdent("Fatal"),
											},
											Args: []ast.Expr{
												&ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"Error in creating postgres instance\"",
												},
											},
										},
									},
								},
							},
						},
						2: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("handler"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("gin"),
										Sel: ast.NewIdent("New"),
									},
								},
							},
						},
						3: &ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("handler"),
									Sel: ast.NewIdent("Use"),
								},
								Args: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   ast.NewIdent("cors"),
											Sel: ast.NewIdent("New"),
										},
										Args: []ast.Expr{
											&ast.CompositeLit{
												Type: &ast.SelectorExpr{
													X:   ast.NewIdent("cors"),
													Sel: ast.NewIdent("Config"),
												},
												Elts: []ast.Expr{
													0: &ast.KeyValueExpr{
														Key: ast.NewIdent("AllowOrigins"),
														Value: &ast.CompositeLit{
															Type: &ast.ArrayType{
																Elt: ast.NewIdent("string"),
															},
															Elts: []ast.Expr{
																&ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "\"*\"",
																},
															},
														},
													},
													1: &ast.KeyValueExpr{
														Key: ast.NewIdent("AllowMethods"),
														Value: &ast.CompositeLit{
															Type: &ast.ArrayType{
																Elt: ast.NewIdent("string"),
															},
															Elts: []ast.Expr{
																&ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "\"*\"",
																},
															},
														},
													},
													2: &ast.KeyValueExpr{
														Key: ast.NewIdent("AllowHeaders"),
														Value: &ast.CompositeLit{
															Type: &ast.ArrayType{
																Elt: ast.NewIdent("string"),
															},
															Elts: []ast.Expr{
																0: &ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "\"Access-Control-Allow-Origin\"",
																},
																1: &ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "\"*\"",
																},
															},
														},
													},
													3: &ast.KeyValueExpr{
														Key: ast.NewIdent("ExposeHeaders"),
														Value: &ast.CompositeLit{
															Type: &ast.ArrayType{
																Elt: ast.NewIdent("string"),
															},
															Elts: []ast.Expr{
																&ast.BasicLit{
																	Kind:  token.STRING,
																	Value: "\"Content-Length\"",
																},
															},
														},
													},
													4: &ast.KeyValueExpr{
														Key:   ast.NewIdent("AllowCredentials"),
														Value: ast.NewIdent("true"),
													},
													5: &ast.KeyValueExpr{
														Key: ast.NewIdent("MaxAge"),
														Value: &ast.BinaryExpr{
															X: &ast.BasicLit{
																Kind:  token.INT,
																Value: "12",
															},
															Op: token.MUL,
															Y: &ast.SelectorExpr{
																X:   ast.NewIdent("time"),
																Sel: ast.NewIdent("Hour"),
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
						4: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("serv"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("httpserver"),
										Sel: ast.NewIdent("New"),
									},
									Args: []ast.Expr{
										0: ast.NewIdent("handler"),
										1: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("httpserver"),
												Sel: ast.NewIdent("Port"),
											},
											Args: []ast.Expr{
												&ast.SelectorExpr{
													X:   ast.NewIdent("cfg"),
													Sel: ast.NewIdent("AppPort"),
												},
											},
										},
									},
								},
							},
						},
						5: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("interruption"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: ast.NewIdent("make"),
									Args: []ast.Expr{
										0: &ast.ChanType{
											Dir: 3,
											Value: &ast.SelectorExpr{
												X:   ast.NewIdent("os"),
												Sel: ast.NewIdent("Signal"),
											},
										},
										1: &ast.BasicLit{
											Kind:  token.INT,
											Value: "1",
										},
									},
								},
							},
						},
						6: &ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("signal"),
									Sel: ast.NewIdent("Notify"),
								},
								Args: []ast.Expr{
									0: ast.NewIdent("interruption"),
									1: &ast.SelectorExpr{
										X:   ast.NewIdent("os"),
										Sel: ast.NewIdent("Interrupt"),
									},
									2: &ast.SelectorExpr{
										X:   ast.NewIdent("syscall"),
										Sel: ast.NewIdent("SIGTERM"),
									},
								},
							},
						},
						7: &ast.SelectStmt{
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									0: &ast.CommClause{
										Comm: &ast.AssignStmt{
											Lhs: []ast.Expr{
												ast.NewIdent("s"),
											},
											Tok: token.DEFINE,
											Rhs: []ast.Expr{
												&ast.UnaryExpr{
													Op: token.ARROW,
													X:  ast.NewIdent("interruption"),
												},
											},
										},
										Body: []ast.Stmt{
											&ast.ExprStmt{
												X: &ast.CallExpr{
													Fun: &ast.SelectorExpr{
														X:   ast.NewIdent("log"),
														Sel: ast.NewIdent("Printf"),
													},
													Args: []ast.Expr{
														0: &ast.BinaryExpr{
															X: &ast.BasicLit{
																Kind:  token.STRING,
																Value: "\"signal: \"",
															},
															Op: token.ADD,
															Y: &ast.CallExpr{
																Fun: &ast.SelectorExpr{
																	X:   ast.NewIdent("s"),
																	Sel: ast.NewIdent("String"),
																},
															},
														},
													},
												},
											},
										},
									},
									1: &ast.CommClause{
										Comm: &ast.AssignStmt{
											Lhs: []ast.Expr{
												ast.NewIdent("err"),
											},
											Tok: token.ASSIGN,
											Rhs: []ast.Expr{
												&ast.UnaryExpr{
													Op: token.ARROW,
													X: &ast.CallExpr{
														Fun: &ast.SelectorExpr{
															X:   ast.NewIdent("serv"),
															Sel: ast.NewIdent("Notify"),
														},
													},
												},
											},
										},
										Body: []ast.Stmt{
											&ast.ExprStmt{
												X: &ast.CallExpr{
													Fun: &ast.SelectorExpr{
														X:   ast.NewIdent("log"),
														Sel: ast.NewIdent("Printf"),
													},
													Args: []ast.Expr{
														&ast.BasicLit{
															Kind:  token.STRING,
															Value: "\"Notify from http server\"",
														},
													},
												},
											},
										},
									},
								},
							},
						},
						8: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("err"),
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("serv"),
										Sel: ast.NewIdent("Shutdown"),
									},
								},
							},
						},
						9: &ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X:  ast.NewIdent("err"),
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("log"),
												Sel: ast.NewIdent("Printf"),
											},
											Args: []ast.Expr{
												&ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"Http server shutdown\"",
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
		},
	}
}
