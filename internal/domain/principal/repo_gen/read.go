package repo_gen

import (
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

func (rg *RepositoryGenerator) genReadRepoAST() *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent(string(strings.ToLower(rg.entAST.Name)[0])),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(rg.entAST.Name + "Repo"),
					},
				},
			},
		},
		Name: ast.NewIdent("Get" + rg.entAST.Name + "By" + rg.entAST.Fields[0].Name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Names: []*ast.Ident{
							ast.NewIdent("ctx"),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
					1: {
						Names: []*ast.Ident{
							ast.NewIdent(strings.ToLower(rg.entAST.Fields[0].Name)),
						},
						Type: ast.NewIdent(util.TypesMap[rg.entAST.Fields[0].TokType]),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(rg.entAST.Name),
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
				0: rg.genQueryAssignRead(),
				1: rg.genExecuteQueryRead(),
				2: rg.genCheckErrorQueryRead(),
				3: rg.genDeferRead(),
				4: rg.genCheckExistsEntityRead(),
				5: rg.genDeclFieldRead(),
				6: rg.genScanResultsQueryRead(),
				7: rg.genCheckErrorScanRead(),
				8: rg.genReturnStatementRead(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genQueryAssignRead() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("query"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: rg.getQueryForRead(),
			},
		},
	}
}

func (rg *RepositoryGenerator) getQueryForRead() string {
	query := strings.Join([]string{
		"SELECT",
		"*",
		"FROM",
		strings.ToLower(rg.entAST.Name), "WHERE",
		strings.ToLower(rg.entAST.Fields[0].Name), "=", "$1"},
		" ")
	return strings.Join([]string{"\"", query, "\""}, "")
}

func (rg *RepositoryGenerator) genExecuteQueryRead() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent("rows"),
			1: ast.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(string(strings.ToLower(rg.entAST.Name)[0])),
							Sel: ast.NewIdent("pg"),
						},
						Sel: ast.NewIdent("Pool"),
					},
					Sel: ast.NewIdent("Query"),
				},
				Args: []ast.Expr{
					0: ast.NewIdent("ctx"),
					1: ast.NewIdent("query"),
					2: ast.NewIdent(strings.ToLower(rg.entAST.Fields[0].Name)),
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genCheckErrorQueryRead() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						0: &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X:   ast.NewIdent("entity"),
								Sel: ast.NewIdent(rg.entAST.Name),
							},
						},
						1: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fmt"),
								Sel: ast.NewIdent("Errorf"),
							},
							Args: []ast.Expr{
								0: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "\"cannot execute query: %w\"",
								},
								1: ast.NewIdent("err"),
							},
						},
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genDeferRead() *ast.DeferStmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Close"),
			},
		},
	}
}

func (rg *RepositoryGenerator) genCheckExistsEntityRead() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.UnaryExpr{
			Op: token.NOT,
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("rows"),
					Sel: ast.NewIdent("Next"),
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						0: &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X:   ast.NewIdent("entity"),
								Sel: ast.NewIdent(rg.entAST.Name),
							},
						},
						1: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fmt"),
								Sel: ast.NewIdent("Errorf"),
							},
							Args: []ast.Expr{
								0: &ast.BasicLit{
									Kind:  token.STRING,
									Value: strings.Join([]string{"\"there is no book with this", strings.ToLower(rg.entAST.Fields[0].Name), ": %v\""}, " "),
								},
								1: ast.NewIdent(strings.ToLower(rg.entAST.Fields[0].Name)),
							},
						},
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genDeclFieldRead() *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						ast.NewIdent(strings.ToLower(rg.entAST.Name)),
					},
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("entity"),
						Sel: ast.NewIdent(rg.entAST.Name),
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genScanResultsQueryRead() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("err"),
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("rows"),
					Sel: ast.NewIdent("Scan"),
				},
				Args: func() []ast.Expr {
					listExpr := make([]ast.Expr, 0)
					for _, field := range rg.entAST.Fields {
						listExpr = append(listExpr, &ast.UnaryExpr{
							Op: token.AND,
							X: &ast.SelectorExpr{
								X:   ast.NewIdent(strings.ToLower(rg.entAST.Name)),
								Sel: ast.NewIdent(field.Name),
							},
						})
					}
					return listExpr
				}(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genCheckErrorScanRead() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						0: &ast.CompositeLit{
							Type: &ast.SelectorExpr{
								X:   ast.NewIdent("entity"),
								Sel: ast.NewIdent(rg.entAST.Name),
							},
						},
						1: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fmt"),
								Sel: ast.NewIdent("Errorf"),
							},
							Args: []ast.Expr{
								0: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "\"error parsing book: %w\"",
								},
								1: ast.NewIdent("err"),
							},
						},
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genReturnStatementRead() *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(rg.entAST.Name)),
			1: ast.NewIdent("nil"),
		},
	}
}
