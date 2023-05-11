package controller_gen

import (
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genCreateControllerAST() *ast.FuncDecl {
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
		Name: ast.NewIdent("create" + cg.entAST.Name),
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
				0: cg.genDefineRequestCreate(),
				1: cg.genParseRequestCreate(),
				2: cg.genCallServiceCreate(),
				3: cg.genCheckCallServiceCreate(),
				4: cg.genHeaderCreate(),
				5: cg.genSendResponseCreate(),
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genDefineRequestCreate() *ast.AssignStmt {
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

func (cg *ControllerHTTPGenerator) genParseRequestCreate() *ast.IfStmt {
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

func (cg *ControllerHTTPGenerator) genCallServiceCreate() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(cg.entAST.Fields[0].Name)),
			1: ast.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
						Sel: ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
					},
					Sel: ast.NewIdent("Create" + cg.entAST.Name),
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
							for i := 1; i < len(cg.entAST.Fields); i++ {
								res = append(res, &ast.KeyValueExpr{
									Key: ast.NewIdent(cg.entAST.Fields[i].Name),
									Value: &ast.SelectorExpr{
										X:   ast.NewIdent("req"),
										Sel: ast.NewIdent(cg.entAST.Fields[i].Name),
									},
								})
							}
							return res
						}(),
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckCallServiceCreate() *ast.IfStmt {
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
								Sel: ast.NewIdent("StatusInternalServerError"),
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

func (cg *ControllerHTTPGenerator) genHeaderCreate() *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("c"),
				Sel: ast.NewIdent("Header"),
			},
			Args: []ast.Expr{
				0: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"Location\"",
				},
				1: &ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("fmt"),
						Sel: ast.NewIdent("Sprintf"),
					},
					Args: []ast.Expr{
						0: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"/api/v1/" + strings.ToLower(cg.entAST.Name) + "/%s\"",
						},
						1: &ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("strconv"),
								Sel: ast.NewIdent("Itoa"),
							},
							Args: []ast.Expr{
								ast.NewIdent(strings.ToLower(cg.entAST.Fields[0].Name)),
							},
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genSendResponseCreate() *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("c"),
				Sel: ast.NewIdent("JSON"),
			},
			Args: []ast.Expr{
				0: &ast.SelectorExpr{
					X:   ast.NewIdent("http"),
					Sel: ast.NewIdent("StatusCreated"),
				},
				1: ast.NewIdent("nil"),
			},
		},
	}
}
