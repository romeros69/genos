package repo_gen

import (
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

func (rg *RepositoryGenerator) genDeleteRepoAST() *ast.FuncDecl {
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
		Name: ast.NewIdent(strings.Join([]string{"Delete", rg.entAST.Name}, "")),
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
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: rg.genQueryAssignDelete(),
				1: rg.genExecuteQueryDelete(),
				2: rg.genCheckErrorQueryDelete(),
				3: rg.genDeferDelete(),
				4: rg.genReturnStatementDelete(),
			},
		},
	}
}

func (rg *RepositoryGenerator) genQueryAssignDelete() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("query"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.BasicLit{
				Kind:  token.STRING,
				Value: rg.getQueryForDelete(),
			},
		},
	}
}

func (rg *RepositoryGenerator) getQueryForDelete() string {
	query := strings.Join([]string{
		"DELETE",
		"*",
		"FROM",
		strings.ToLower(rg.entAST.Name),
		"WHERE",
		strings.ToLower(rg.entAST.Fields[0].Name),
		"=",
		"$1"}, " ")
	return strings.Join([]string{"\"", query, "\""}, "")
}

func (rg *RepositoryGenerator) genExecuteQueryDelete() *ast.AssignStmt {
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

func (rg *RepositoryGenerator) genCheckErrorQueryDelete() *ast.IfStmt {
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

func (rg *RepositoryGenerator) genDeferDelete() *ast.DeferStmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("rows"),
				Sel: ast.NewIdent("Close"),
			},
		},
	}
}

func (rg *RepositoryGenerator) genReturnStatementDelete() *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: []ast.Expr{
			ast.NewIdent("nil"),
		},
	}
}
