package repo_gen

import (
	"go/ast"
	"go/token"
	"strings"
)

func (rg *RepositoryGenerator) genListRepoAST() *ast.FuncDecl {
	nameEntity := rg.entAST.Name
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent(string(strings.ToLower(nameEntity)[0])),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(nameEntity + "Repo"),
					},
				},
			},
		},
		Name: ast.NewIdent("Get" + nameEntity + "List"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("ctx"),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.ArrayType{
							Elt: &ast.SelectorExpr{
								X:   ast.NewIdent("entity"),
								Sel: ast.NewIdent(nameEntity),
							},
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
				0: rg.genQueryAssignList(),
				1: rg.genExecuteQueryList(),
				2: rg.genCheckErrorQueryList(),
				3: rg.genDeferList(),
				4: rg.genDeclListEntityList(),
				5: rg.genCycleForReadResultsQueryList(),
				6: rg.genReturnStatementList(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genQueryAssignList() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("query"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: rg.getQueryForList(),
			},
		},
	}
}

func (rg *RepositoryGenerator) getQueryForList() string {
	query := strings.Join([]string{
		"SELECT",
		"*",
		"FROM",
		strings.ToLower(rg.entAST.Name),
	}, " ")
	return strings.Join([]string{"\"", query, "\""}, "")
}

func (rg *RepositoryGenerator) genExecuteQueryList() *ast.AssignStmt {
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
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genCheckErrorQueryList() *ast.IfStmt {
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
						0: ast.NewIdent("nil"),
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

func (rg *RepositoryGenerator) genDeferList() *ast.DeferStmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Close"),
			},
		},
	}
}

func (rg *RepositoryGenerator) genDeclListEntityList() *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						ast.NewIdent(strings.ToLower(rg.entAST.Name) + "List"),
					},
					Type: &ast.ArrayType{
						Elt: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(rg.entAST.Name),
						},
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genCycleForReadResultsQueryList() *ast.ForStmt {
	return &ast.ForStmt{
		Cond: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Next"),
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: &ast.DeclStmt{
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
				},
				1: &ast.AssignStmt{
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
							Args: rg.generateExprListForScanList(),
						},
					},
				},
				2: &ast.IfStmt{
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
												Value: "\"error in parsing book: %w\"",
											},
											1: ast.NewIdent("err"),
										},
									},
								},
							},
						},
					},
				},
				3: &ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent(strings.ToLower(rg.entAST.Name) + "List"),
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: ast.NewIdent("append"),
							Args: []ast.Expr{
								0: ast.NewIdent(strings.ToLower(rg.entAST.Name) + "List"),
								1: ast.NewIdent(strings.ToLower(rg.entAST.Name)),
							},
						},
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genReturnStatementList() *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(rg.entAST.Name) + "List"),
			1: ast.NewIdent("nil"),
		},
	}
}

// функция генерация слайса выражений для считвания результата запроса list
func (rg *RepositoryGenerator) generateExprListForScanList() []ast.Expr {
	eSlice := make([]ast.Expr, len(rg.entAST.Fields))
	for i, v := range rg.entAST.Fields {
		eSlice[i] = &ast.UnaryExpr{
			Op: token.AND,
			X: &ast.SelectorExpr{
				X:   ast.NewIdent(strings.ToLower(rg.entAST.Name)),
				Sel: ast.NewIdent(v.Name),
			},
		}
	}
	return eSlice
}
