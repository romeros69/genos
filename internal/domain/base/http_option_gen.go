package base

import (
	"go/ast"
	"go/token"
)

type HttpOptionsGenerator struct {
	fullPathToFile string
}

var _ BaseGenerator = (*HttpOptionsGenerator)(nil)

func NewHttpOptionsGenerator() *HttpOptionsGenerator {
	return &HttpOptionsGenerator{
		fullPathToFile: "pkg/httpserver/options.go",
	}
}

func (ho *HttpOptionsGenerator) FullPathToFile() string {
	return ho.fullPathToFile
}

func (ho *HttpOptionsGenerator) GenAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("httpserver"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"net\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"time\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Option"),
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{
										Type: &ast.StarExpr{
											X: ast.NewIdent("Server"),
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("Port"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("port"),
								},
								Type: ast.NewIdent("string"),
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("Addr"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													&ast.CallExpr{
														Fun: &ast.SelectorExpr{
															X:   ast.NewIdent("net"),
															Sel: ast.NewIdent("JoinHostPort"),
														},
														Args: []ast.Expr{
															0: &ast.BasicLit{
																Kind:  token.STRING,
																Value: "\"\"",
															},
															1: ast.NewIdent("port"),
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
			},
			// 3
			3: &ast.FuncDecl{
				Name: ast.NewIdent("ReadTimeout"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("timeout"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Duration"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("ReadTimeout"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("timeout"),
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
			4: &ast.FuncDecl{
				Name: ast.NewIdent("WriteTimeout"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("timeout"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Duration"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("WriteTimeout"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("timeout"),
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
			5: &ast.FuncDecl{
				Name: ast.NewIdent("ShutdownTimeout"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("timeout"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Duration"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X:   ast.NewIdent("s"),
														Sel: ast.NewIdent("shutdownTimeout"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("timeout"),
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
