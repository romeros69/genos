package repo_gen

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
)

func (rg *RepositoryGenerator) genUpdateRepoAST() *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent(string(strings.ToLower(rg.entAST.Name)[0])),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(strings.Join([]string{rg.entAST.Name, "Repo"}, "")),
					},
				},
			},
		},
		Name: ast.NewIdent(strings.Join([]string{"Update", rg.entAST.Name}, "")),
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
							ast.NewIdent(strings.ToLower(rg.entAST.Name)),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(rg.entAST.Name),
						},
					},
				},
			},
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
				0: rg.genQueryAssignUpdate(),
				1: rg.genExecuteQueryUpdate(),
				2: rg.genCheckErrorQueryUpdate(),
				3: rg.genDeferUpdate(),
				4: rg.genReturnStatementUpdate(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genQueryAssignUpdate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("query"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: rg.getQueryForUpdate(),
			},
		},
	}
}

func (rg *RepositoryGenerator) getQueryForUpdate() string {
	strFields := ""
	for i := 1; i < len(rg.entAST.Fields); i++ {
		if i != len(rg.entAST.Fields)-1 {
			strFields += strings.Join([]string{strings.ToLower(rg.entAST.Fields[i].Name), "=$", strconv.Itoa(i), ", "}, "")
		} else {
			strFields += strings.Join([]string{strings.ToLower(rg.entAST.Fields[i].Name), "=$", strconv.Itoa(i)}, "")
		}
	}
	query := strings.Join([]string{
		"UPDATE",
		strings.ToLower(rg.entAST.Name),
		"SET",
		strFields,
		"WHERE",
		strings.ToLower(rg.entAST.Fields[0].Name), "=", "$" + strconv.Itoa(len(rg.entAST.Fields))}, " ")
	return strings.Join([]string{"\"", query, "\""}, "")
}

func (rg *RepositoryGenerator) genExecuteQueryUpdate() *ast.AssignStmt {
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
				Args: func() []ast.Expr {
					result := []ast.Expr{
						0: ast.NewIdent("ctx"),
						1: ast.NewIdent("query"),
					}
					for i := 1; i < len(rg.entAST.Fields); i++ {
						result = append(result, &ast.SelectorExpr{
							X:   ast.NewIdent(strings.ToLower(rg.entAST.Name)),
							Sel: ast.NewIdent(rg.entAST.Fields[i].Name),
						})
					}
					return result
				}(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genCheckErrorQueryUpdate() *ast.IfStmt {
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
						&ast.CallExpr{
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

func (rg *RepositoryGenerator) genDeferUpdate() *ast.DeferStmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Close"),
			},
		},
	}
}

func (rg *RepositoryGenerator) genReturnStatementUpdate() *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: []ast.Expr{
			ast.NewIdent("nil"),
		},
	}
}
