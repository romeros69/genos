package base

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func createPostgresOptionsAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("postgres"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
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
											X: ast.NewIdent("Postgres"),
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("MaxPoolSize"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("size"),
								},
								Type: ast.NewIdent("int"),
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
														ast.NewIdent("c"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Postgres"),
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
														X:   ast.NewIdent("c"),
														Sel: ast.NewIdent("maxPoolSize"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("size"),
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
			3: &ast.FuncDecl{
				Name: ast.NewIdent("ConnAttempts"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("attempts"),
								},
								Type: ast.NewIdent("int"),
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
														ast.NewIdent("c"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Postgres"),
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
														X:   ast.NewIdent("c"),
														Sel: ast.NewIdent("connAttempts"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("attempts"),
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
				Name: ast.NewIdent("ConnTimeout"),
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
														ast.NewIdent("c"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Postgres"),
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
														X:   ast.NewIdent("c"),
														Sel: ast.NewIdent("connTimeout"),
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

func GenPostgresOptions() error {
	f := createPostgresOptionsAST()
	fset := token.NewFileSet()

	file, err := os.Create("pkg/postgres/options.go")
	if err != nil {
		return fmt.Errorf("error in creating main.go file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error in closing file %s", file.Name())
		}
	}(file)

	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in genereate postgres option: %w", err)
	}
	return nil
}
