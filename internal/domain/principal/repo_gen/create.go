package repo_gen

import (
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

func (rg *RepositoryGenerator) genCreateRepoAST() *ast.FuncDecl {
	nameEntity := rg.entAST.Name
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent(string(strings.ToLower(nameEntity)[0]))},
					Type: &ast.StarExpr{
						X: ast.NewIdent(nameEntity + "Repo"),
					},
				},
			},
		},
		Name: ast.NewIdent("Create" + nameEntity),
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
						Names: []*ast.Ident{ast.NewIdent(strings.ToLower(nameEntity))},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(nameEntity),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: ast.NewIdent(util.TypesMap[rg.entAST.Fields[0].TokType]),
					},
					1: {
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: rg.genQueryAssignCreate(),
				1: rg.genExecuteQueryCreate(),
				2: rg.genCheckErrorQueryCreate(),
				3: rg.genDeferCreate(),
				4: rg.genDeclFieldCreate(),
				5: rg.genCycleForReadResultsQueryCreate(),
				6: rg.genReturnStatementCreate(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genQueryAssignCreate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("query"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: rg.getQueryForCreate(),
			},
		},
	}
}

// getQueryForCreate - собирает строку запроса для операции CREATE
func (rg *RepositoryGenerator) getQueryForCreate() string {
	listNamesFields := ""
	argsList := ""
	for i, field := range rg.entAST.Fields {
		if i != 0 {
			if i == len(rg.entAST.Fields)-1 {
				listNamesFields += field.Name
				argsList += "$" + strconv.Itoa(i)
			} else {
				listNamesFields += field.Name + ", "
				argsList += "$" + strconv.Itoa(i) + ", "
			}
		}
	}
	listNamesFields = strings.Join([]string{"(", listNamesFields, ")"}, "")
	argsList = strings.Join([]string{"(", argsList, ")"}, "")
	query := strings.Join([]string{"INSERT INTO",
		strings.ToLower(rg.entAST.Name),
		listNamesFields,
		"VALUES",
		argsList,
		"RETURNING",
		strings.ToLower(rg.entAST.Fields[0].Name)}, " ")
	return strings.Join([]string{"\"", query, "\""}, "")
}

func (rg *RepositoryGenerator) genExecuteQueryCreate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent("rows"),
			ast.NewIdent("err"),
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
				Args: func() []ast.Expr {
					res := []ast.Expr{
						0: ast.NewIdent("ctx"),
						1: ast.NewIdent("query"),
					}
					for i := 1; i < len(rg.entAST.Fields); i++ {
						res = append(res, &ast.SelectorExpr{
							X:   ast.NewIdent(strings.ToLower(rg.entAST.Name)),
							Sel: ast.NewIdent(rg.entAST.Fields[i].Name),
						})
					}
					return res
				}(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genCheckErrorQueryCreate() *ast.IfStmt {
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
						0: func() ast.Expr {
							var tokVal token.Token
							if tokVal = util.GetGoTokenByLexerToken(rg.entAST.Fields[0].TokType); tokVal != -1 {
								return &ast.BasicLit{
									Kind:  tokVal,
									Value: util.GetDefaultValueOfType(util.TypesMap[rg.entAST.Fields[0].TokType]),
								}
							}
							return ast.NewIdent("false")
						}(),
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

func (rg *RepositoryGenerator) genDeferCreate() *ast.DeferStmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Close"),
			},
		},
	}
}

func (rg *RepositoryGenerator) genDeclFieldCreate() *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						ast.NewIdent(strings.ToLower(rg.entAST.Fields[0].Name)),
					},
					Type: ast.NewIdent(util.TypesMap[rg.entAST.Fields[0].TokType]),
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) genCycleForReadResultsQueryCreate() *ast.ForStmt {
	return &ast.ForStmt{
		Cond: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Next"),
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: &ast.AssignStmt{
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
							Args: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X:  ast.NewIdent(strings.ToLower(rg.entAST.Fields[0].Name)),
								},
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
							&ast.ReturnStmt{
								Results: []ast.Expr{
									0: func() ast.Expr {
										var tokVal token.Token
										if tokVal = util.GetGoTokenByLexerToken(rg.entAST.Fields[0].TokType); tokVal != -1 {
											return &ast.BasicLit{
												Kind:  tokVal,
												Value: util.GetDefaultValueOfType(util.TypesMap[rg.entAST.Fields[0].TokType]),
											}
										}
										return ast.NewIdent("false")
									}(),
									1: &ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   ast.NewIdent("fmt"),
											Sel: ast.NewIdent("Errorf"),
										},
										Args: []ast.Expr{
											0: &ast.BasicLit{
												Kind:  token.STRING,
												Value: "\"error parsing id book: %w\"",
											},
											1: ast.NewIdent("err"),
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

func (rg *RepositoryGenerator) genReturnStatementCreate() *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(rg.entAST.Fields[0].Name)),
			1: ast.NewIdent("err"),
		},
	}
}
