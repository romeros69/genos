package controller_gen

import (
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genUpdateControllerAST() *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Routes"),
					},
				},
			},
		},
		Name: ast.NewIdent("update" + cg.entAST.Name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							ast.NewIdent("c"),
						},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("gin"),
								Sel: ast.NewIdent("Context"),
							},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: cg.genGetQueryParamUpdate(),
				1: cg.genCheckGetQueryParamUpdate(),
				2: cg.genDefineRequestUpdate(),
				3: cg.genParseRequestUpdate(),
				4: cg.genCallServiceUpdate(),
				5: cg.genCheckCallServiceUpdate(),
				6: cg.genSendResponseUpdate(),
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genGetQueryParamUpdate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "ID"),
			1: ast.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("strconv"),
					Sel: ast.NewIdent("Atoi"),
				},
				Args: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("c"),
							Sel: ast.NewIdent("Param"),
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"" + strings.ToLower(cg.entAST.Fields[0].Name) + "\"",
							},
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckGetQueryParamUpdate() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: ast.NewIdent("errorResponse"),
						Args: []ast.Expr{
							0: ast.NewIdent("c"),
							1: &ast.SelectorExpr{
								X:   ast.NewIdent("http"),
								Sel: ast.NewIdent("StatusBadRequest"),
							},
							2: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("err"),
									Sel: ast.NewIdent("Error"),
								},
							},
						},
					},
				},
				1: &ast.ReturnStmt{},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genDefineRequestUpdate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("req"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: ast.NewIdent("new"),
				Args: []ast.Expr{
					ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Request"),
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genParseRequestUpdate() *ast.IfStmt {
	return &ast.IfStmt{
		Init: &ast.AssignStmt{
			Lhs: []ast.Expr{
				ast.NewIdent("err"),
			},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("c"),
						Sel: ast.NewIdent("ShouldBindJSON"),
					},
					Args: []ast.Expr{
						ast.NewIdent("req"),
					},
				},
			},
		},
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: ast.NewIdent("errorResponse"),
						Args: []ast.Expr{
							0: ast.NewIdent("c"),
							1: &ast.SelectorExpr{
								X:   ast.NewIdent("http"),
								Sel: ast.NewIdent("StatusBadRequest"),
							},
							2: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("err"),
									Sel: ast.NewIdent("Error"),
								},
							},
						},
					},
				},
				1: &ast.ReturnStmt{},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCallServiceUpdate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			ast.NewIdent("err"),
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
						Sel: ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
					},
					Sel: ast.NewIdent("Update" + cg.entAST.Name),
				},
				Args: []ast.Expr{
					0: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("c"),
								Sel: ast.NewIdent("Request"),
							},
							Sel: ast.NewIdent("Context"),
						},
					},
					1: &ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(cg.entAST.Name),
						},
						Elts: func() []ast.Expr {
							res := make([]ast.Expr, 0)
							for i, v := range cg.entAST.Fields {
								if i == 0 {
									res = append(res, &ast.KeyValueExpr{
										Key: ast.NewIdent(v.Name),
										Value: &ast.CallExpr{
											Fun: ast.NewIdent(util.TypesMap[cg.entAST.Fields[0].TokType]),
											Args: []ast.Expr{
												ast.NewIdent(strings.ToLower(cg.entAST.Name) + v.Name),
											},
										},
									})
								} else {
									res = append(res, &ast.KeyValueExpr{
										Key: ast.NewIdent(v.Name),
										Value: &ast.SelectorExpr{
											X:   ast.NewIdent("req"),
											Sel: ast.NewIdent(v.Name),
										},
									})
								}
							}
							return res
						}(),
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckCallServiceUpdate() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				0: &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: ast.NewIdent("errorResponse"),
						Args: []ast.Expr{
							0: ast.NewIdent("c"),
							1: &ast.SelectorExpr{
								X:   ast.NewIdent("http"),
								Sel: ast.NewIdent("StatusNotFound"),
							},
							2: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("err"),
									Sel: ast.NewIdent("Error"),
								},
							},
						},
					},
				},
				1: &ast.ReturnStmt{},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genSendResponseUpdate() *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("c"),
				Sel: ast.NewIdent("JSON"),
			},
			Args: []ast.Expr{
				0: &ast.SelectorExpr{
					X:   ast.NewIdent("http"),
					Sel: ast.NewIdent("StatusNoContent"),
				},
				1: ast.NewIdent("nil"),
			},
		},
	}
}
