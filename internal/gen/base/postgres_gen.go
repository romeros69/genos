package base

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func createPostgresAST(moduleName string) *ast.File {
	return &ast.File{
		Name: ast.NewIdent("postgres"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"context\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
					2: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"log\"",
						},
					},
					3: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/configs\"",
						},
					},
					4: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"time\"",
						},
					},
					5: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/Masterminds/squirrel\"",
						},
					},
					6: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/jackc/pgx/v4/pgxpool\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.CONST,
				Specs: []ast.Spec{
					0: &ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultMaxPoolSize"),
						},
						Values: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.INT,
								Value: "1",
							},
						},
					},
					1: &ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultConnAttempts"),
						},
						Values: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.INT,
								Value: "10",
							},
						},
					},
					2: &ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_defaultConnTimeout"),
						},
						Values: []ast.Expr{
							&ast.SelectorExpr{
								X:   ast.NewIdent("time"),
								Sel: ast.NewIdent("Second"),
							},
						},
					},
				},
			},
			2: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Postgres"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									0: {
										Names: []*ast.Ident{
											ast.NewIdent("maxPoolSize"),
										},
										Type: ast.NewIdent("int32"),
									},
									1: {
										Names: []*ast.Ident{
											ast.NewIdent("connAttempts"),
										},
										Type: ast.NewIdent("int"),
									},
									2: {
										Names: []*ast.Ident{
											ast.NewIdent("connTimeout"),
										},
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("time"),
											Sel: ast.NewIdent("Duration"),
										},
									},
									3: {
										Names: []*ast.Ident{
											ast.NewIdent("Builder"),
										},
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("squirrel"),
											Sel: ast.NewIdent("StatementBuilderType"),
										},
									},
									4: {
										Names: []*ast.Ident{
											ast.NewIdent("Pool"),
										},
										Type: &ast.StarExpr{
											X: &ast.SelectorExpr{
												X:   ast.NewIdent("pgxpool"),
												Sel: ast.NewIdent("Pool"),
											},
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
					Results: &ast.FieldList{
						List: []*ast.Field{
							0: {
								Type: &ast.StarExpr{
									X: ast.NewIdent("Postgres"),
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
						0: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("pg"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: ast.NewIdent("Postgres"),
										Elts: []ast.Expr{
											0: &ast.KeyValueExpr{
												Key:   ast.NewIdent("maxPoolSize"),
												Value: ast.NewIdent("_defaultMaxPoolSize"),
											},
											1: &ast.KeyValueExpr{
												Key:   ast.NewIdent("connAttempts"),
												Value: ast.NewIdent("_defaultConnAttempts"),
											},
											2: &ast.KeyValueExpr{
												Key:   ast.NewIdent("connTimeout"),
												Value: ast.NewIdent("_defaultConnTimeout"),
											},
										},
									},
								},
							},
						},
						1: &ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("pg"),
									Sel: ast.NewIdent("Builder"),
								},
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X: &ast.SelectorExpr{
											X:   ast.NewIdent("squirrel"),
											Sel: ast.NewIdent("StatementBuilder"),
										},
										Sel: ast.NewIdent("PlaceholderFormat"),
									},
									Args: []ast.Expr{
										&ast.SelectorExpr{
											X:   ast.NewIdent("squirrel"),
											Sel: ast.NewIdent("Dollar"),
										},
									},
								},
							},
						},
						2: &ast.AssignStmt{
							Lhs: []ast.Expr{
								0: ast.NewIdent("poolConfig"),
								1: ast.NewIdent("err"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("pgxpool"),
										Sel: ast.NewIdent("ParseConfig"),
									},
									Args: []ast.Expr{
										&ast.SelectorExpr{
											X:   ast.NewIdent("cfg"),
											Sel: ast.NewIdent("PostgresURL"),
										},
									},
								},
							},
						},
						3: &ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X:  ast.NewIdent("err"),
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ReturnStmt{
										Results: []ast.Expr{
											0: ast.NewIdent("nil"),
											1: &ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X:   ast.NewIdent("fmt"),
													Sel: ast.NewIdent("Errorf"),
												},
												Args: []ast.Expr{
													0: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "\"postgres - NewPostgres - pgxpool.ParseConfig: %w\"",
													},
													1: ast.NewIdent("err"),
												},
											},
										},
									},
								},
							},
						},
						4: &ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("poolConfig"),
									Sel: ast.NewIdent("MaxConns"),
								},
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("pg"),
									Sel: ast.NewIdent("maxPoolSize"),
								},
							},
						},
						5: &ast.ForStmt{
							Cond: &ast.BinaryExpr{
								X: &ast.SelectorExpr{
									X:   ast.NewIdent("pg"),
									Sel: ast.NewIdent("connAttempts"),
								},
								Op: token.GTR,
								Y: &ast.BasicLit{
									Kind:  token.INT,
									Value: "0",
								},
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									0: &ast.AssignStmt{
										Lhs: []ast.Expr{
											0: &ast.SelectorExpr{
												X:   ast.NewIdent("pg"),
												Sel: ast.NewIdent("Pool"),
											},
											1: ast.NewIdent("err"),
										},
										Tok: token.ASSIGN,
										Rhs: []ast.Expr{
											&ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X:   ast.NewIdent("pgxpool"),
													Sel: ast.NewIdent("ConnectConfig"),
												},
												Args: []ast.Expr{
													0: &ast.CallExpr{
														Fun: &ast.SelectorExpr{
															X:   ast.NewIdent("context"),
															Sel: ast.NewIdent("Background"),
														},
													},
													1: ast.NewIdent("poolConfig"),
												},
											},
										},
									},
									1: &ast.IfStmt{
										Cond: &ast.BinaryExpr{
											X:  ast.NewIdent("err"),
											Op: token.EQL,
											Y:  ast.NewIdent("nil"),
										},
										Body: &ast.BlockStmt{
											List: []ast.Stmt{
												&ast.BranchStmt{
													Tok: token.BREAK,
												},
											},
										},
									},
									2: &ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("log"),
												Sel: ast.NewIdent("Printf"),
											},
											Args: []ast.Expr{
												0: &ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"Postgres is trying to connect, attempts left: %d\"",
												},
												1: &ast.SelectorExpr{
													X:   ast.NewIdent("pg"),
													Sel: ast.NewIdent("connAttempts"),
												},
											},
										},
									},
									3: &ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("time"),
												Sel: ast.NewIdent("Sleep"),
											},
											Args: []ast.Expr{
												&ast.SelectorExpr{
													X:   ast.NewIdent("pg"),
													Sel: ast.NewIdent("connTimeout"),
												},
											},
										},
									},
									4: &ast.IncDecStmt{
										X: &ast.SelectorExpr{
											X:   ast.NewIdent("pg"),
											Sel: ast.NewIdent("connAttempts"),
										},
										Tok: token.DEC,
									},
								},
							},
						},
						6: &ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X:  ast.NewIdent("err"),
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ReturnStmt{
										Results: []ast.Expr{
											0: ast.NewIdent("nil"),
											1: &ast.CallExpr{
												Fun: &ast.SelectorExpr{
													X:   ast.NewIdent("fmt"),
													Sel: ast.NewIdent("Errorf"),
												},
												Args: []ast.Expr{
													0: &ast.BasicLit{
														Kind:  token.STRING,
														Value: "\"postgres - NewPostgres - connAttempts == 0: %w\"",
													},
													1: ast.NewIdent("err"),
												},
											},
										},
									},
								},
							},
						},
						7: &ast.ReturnStmt{
							Results: []ast.Expr{
								0: ast.NewIdent("pg"),
								1: ast.NewIdent("nil"),
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
								ast.NewIdent("p"),
							},
							Type: &ast.StarExpr{
								X: ast.NewIdent("Postgres"),
							},
						},
					},
				},
				Name: ast.NewIdent("Close"),
				Type: &ast.FuncType{
					Params: nil,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X: &ast.SelectorExpr{
									X:   ast.NewIdent("p"),
									Sel: ast.NewIdent("Pool"),
								},
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X: &ast.SelectorExpr{
													X:   ast.NewIdent("p"),
													Sel: ast.NewIdent("Pool"),
												},
												Sel: ast.NewIdent("Close"),
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

func GenPostgres(moduleName string) error {
	f := createPostgresAST(moduleName)
	fset := token.NewFileSet()

	file, err := os.Create("pkg/postgres/postgres.go")
	if err != nil {
		return fmt.Errorf("error in creating postgres.go file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error in closing file %s", file.Name())
		}
	}(file)

	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in genereate postgres: %w", err)
	}
	return nil
}
