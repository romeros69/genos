package generate

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

// GenMain - генерация main.go
func GenMain(nameModule string) error {
	f := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + nameModule + "/internal/" + "app\"",
						},
					},
				},
			},
			&ast.FuncDecl{
				Name: ast.NewIdent("main"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						Opening: 10,
						Closing: 11,
					},
				},
				Body: &ast.BlockStmt{
					//Lbrace: 13,
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("app"),
									Sel: ast.NewIdent("Run"),
								},
								//Args: []ast.Expr{
								//	&ast.BasicLit{
								//		Kind:  token.STRING,
								//		Value: "\"hello, it is genos\"",
								//	},
								//},
							},
						},
					},
				},
			},
		},
	}
	fset := token.NewFileSet()

	file, err := os.Create("cmd/main/main.go")
	if err != nil {
		return fmt.Errorf("error in creating main.go file: %w", err)
	}
	defer file.Close()

	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in genereate main: %w", err)
	}
	return nil
}

// GenApp - Генерация App.go
func GenApp() error {
	f := &ast.File{
		Name: ast.NewIdent("app"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
				},
			},
			&ast.FuncDecl{
				Name: ast.NewIdent("Run"),
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("fmt"),
									Sel: ast.NewIdent("Println"),
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: "\"Run complete!\"",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	fset := token.NewFileSet()
	file, err := os.Create("internal/app/app.go")
	if err != nil {
		return fmt.Errorf("error creating app.go file: %w", err)
	}
	defer file.Close()
	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in generate app.go: %w", err)
	}
	return nil
}

// GenHttpServer - Генерация server.go
func GenHttpServer() error {
	f := &ast.File{
		Name: ast.NewIdent("httpserver"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"context\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"net/http\"",
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
				Tok: token.CONST,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultReadTimeout"),
						},
						Values: []ast.Expr{
							&ast.BinaryExpr{
								X: &ast.BasicLit{
									Kind:  token.INT,
									Value: "5",
								},
								Op: token.MUL,
								Y: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Second"),
								},
							},
						},
					},
					&ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultWriteTimeout"),
						},
						Values: []ast.Expr{
							&ast.BinaryExpr{
								X: &ast.BasicLit{
									Kind:  token.INT,
									Value: "5",
								},
								Op: token.MUL,
								Y: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Second"),
								},
							},
						},
					},
					&ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultAddr"),
						},
						Values: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\":80\"",
							},
						},
					},
					&ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultShutdownTimeout"),
						},
						Values: []ast.Expr{
							&ast.BinaryExpr{
								X: &ast.BasicLit{
									Kind:  token.INT,
									Value: "3",
								},
								Op: token.MUL,
								Y: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Second"),
								},
							},
						},
					},
				},
			},
			2: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Server"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											0: ast.NewIdent("server"),
										},
										Type: &ast.StarExpr{ // Указатель
											X: &ast.SelectorExpr{
												X:   ast.NewIdent("http"),
												Sel: ast.NewIdent("Server"),
											},
										},
									},
									{
										Names: []*ast.Ident{
											ast.NewIdent("notify"),
										},
										Type: &ast.ChanType{
											Dir:   3, // type of direction chanel
											Value: ast.NewIdent("error"),
										},
									},
									{
										Names: []*ast.Ident{
											0: ast.NewIdent("shutdownTimeout"),
										},
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("time"),
											Sel: ast.NewIdent("Duration"),
										},
									},
								},
							},
						},
					},
				},
			},
			3: &ast.FuncDecl{
				Name: ast.NewIdent("New"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							0: {
								Names: []*ast.Ident{
									ast.NewIdent("handler"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("http"),
									Sel: ast.NewIdent("Handler"),
								},
							},
							1: {
								Names: []*ast.Ident{
									ast.NewIdent("opts"),
								},
								Type: &ast.Ellipsis{
									Elt: ast.NewIdent("Option"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: ast.NewIdent("Server"),
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						0: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("httpServer"),
							},
							Tok: token.DEFINE, // :=
							Rhs: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("http"),
											Sel: ast.NewIdent("Server"),
										},
										Elts: []ast.Expr{
											0: &ast.KeyValueExpr{
												Key:   ast.NewIdent("Handler"),
												Value: ast.NewIdent("handler"),
											},
											1: &ast.KeyValueExpr{
												Key:   ast.NewIdent("ReadTimeout"),
												Value: ast.NewIdent("_defaultReadTimeout"),
											},
											2: &ast.KeyValueExpr{
												Key:   ast.NewIdent("WriteTimeout"),
												Value: ast.NewIdent("_defaultWriteTimeout"),
											},
											3: &ast.KeyValueExpr{
												Key:   ast.NewIdent("Addr"),
												Value: ast.NewIdent("_defaultAddr"),
											},
										},
										Incomplete: false, // check
									},
								},
							},
						},
						1: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("s"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: ast.NewIdent("Server"),
										Elts: []ast.Expr{
											0: &ast.KeyValueExpr{
												Key:   ast.NewIdent("server"),
												Value: ast.NewIdent("httpServer"),
											},
											1: &ast.KeyValueExpr{
												Key: ast.NewIdent("notify"),
												Value: &ast.CallExpr{
													Fun: ast.NewIdent("make"),
													Args: []ast.Expr{
														0: &ast.ChanType{
															Dir:   3,
															Value: ast.NewIdent("error"),
														},
														1: &ast.BasicLit{
															Kind:  token.INT,
															Value: "1",
														},
													},
												},
											},
											2: &ast.KeyValueExpr{
												Key:   ast.NewIdent("shutdownTimeout"),
												Value: ast.NewIdent("_defaultShutdownTimeout"),
											},
										},
									},
								},
							},
						},
						2: &ast.RangeStmt{
							Key: &ast.Ident{
								Name: "_",
								Obj: &ast.Object{
									Kind: ast.Var,
									Name: "_",
									Decl: &ast.AssignStmt{
										Lhs: []ast.Expr{
											0: ast.NewIdent("_"), // Link to +7
											1: ast.NewIdent("opt"),
										},
										Tok: token.DEFINE,
										Rhs: []ast.Expr{
											&ast.UnaryExpr{
												Op: token.RANGE,
												X:  ast.NewIdent("opts"), // link to first decl opts ident
											},
										},
									},
								},
							},
							Value: ast.NewIdent("opt"),
							Tok:   token.DEFINE,
							X:     ast.NewIdent("opts"),
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: ast.NewIdent("opt"),
											Args: []ast.Expr{
												ast.NewIdent("s"),
											},
										},
									},
								},
							},
						},
						3: &ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("s"),
									Sel: ast.NewIdent("start"),
								},
							},
						},
						4: &ast.ReturnStmt{
							Results: []ast.Expr{
								ast.NewIdent("s"),
							},
						},
					},
				},
			},
			4: &ast.FuncDecl{
				Recv: &ast.FieldList{
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
				Name: ast.NewIdent("start"),
				Type: &ast.FuncType{
					Params: nil,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.GoStmt{
							Call: &ast.CallExpr{
								Fun: &ast.FuncLit{
									Type: &ast.FuncType{
										Params: nil,
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											0: &ast.SendStmt{
												Chan: &ast.SelectorExpr{
													X:   ast.NewIdent("s"),
													Sel: ast.NewIdent("notify"),
												},
												Value: &ast.CallExpr{
													Fun: &ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("ListenAndServe"),
													},
												},
											},
											1: &ast.ExprStmt{
												X: &ast.CallExpr{
													Fun: ast.NewIdent("close"),
													Args: []ast.Expr{
														&ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("notify"),
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
			5: &ast.FuncDecl{
				Recv: &ast.FieldList{
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
				Name: ast.NewIdent("Notify"),
				Type: &ast.FuncType{
					Params: nil,
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.ChanType{
									Dir:   2,
									Value: ast.NewIdent("error"),
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("s"),
									Sel: ast.NewIdent("notify"),
								},
							},
						},
					},
				},
			},
			6: &ast.FuncDecl{
				Recv: &ast.FieldList{
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
				Name: ast.NewIdent("Shutdown"),
				Type: &ast.FuncType{
					Params: nil,
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
						0: &ast.AssignStmt{
							Lhs: []ast.Expr{
								0: ast.NewIdent("ctx"),
								1: ast.NewIdent("cancel"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("context"),
										Sel: ast.NewIdent("WithTimeout"),
									},
									Args: []ast.Expr{
										0: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("context"),
												Sel: ast.NewIdent("Background"),
											},
										},
										1: &ast.SelectorExpr{
											X:   ast.NewIdent("s"),
											Sel: ast.NewIdent("shutdownTimeout"),
										},
									},
								},
							},
						},
						1: &ast.DeferStmt{
							Call: &ast.CallExpr{
								Fun: ast.NewIdent("cancel"),
							},
						},
						2: &ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.SelectorExpr{
											X:   ast.NewIdent("s"),
											Sel: ast.NewIdent("server"),
										},
										Sel: ast.NewIdent("Shutdown"),
									},
									Args: []ast.Expr{
										ast.NewIdent("ctx"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	fset := token.NewFileSet()
	file, err := os.Create("pkg/httpserver/server.go")
	if err != nil {
		return fmt.Errorf("error creating app.go file: %w", err)
	}
	defer file.Close()
	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in generate server.go (http server): %w", err)
	}
	return nil
}

// GenOptionsHttpServer - Генерация option.go (httpserver)
func GenOptionsHttpServer() error {

	f := &ast.File{
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

	fset := token.NewFileSet()
	file, err := os.Create("pkg/httpserver/option.go")
	if err != nil {
		return fmt.Errorf("error creating app.go file: %w", err)
	}
	defer file.Close()
	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in generate server.go (http server): %w", err)
	}
	return nil
}
