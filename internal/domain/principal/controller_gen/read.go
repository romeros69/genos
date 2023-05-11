package controller_gen

import (
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genReadControllerAST() *ast.FuncDecl {
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
		Name: ast.NewIdent("get" + cg.entAST.Name + "By" + cg.entAST.Fields[0].Name),
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
				0: cg.genGetParamQueryRead(),
				1: cg.genCheckGetParamRead(),
				2: cg.genCallServiceRead(),
				3: cg.genCheckCallServiceRead(),
				4: cg.genSendResponseRead(),
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genGetParamQueryRead() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "ID"),
			1: ast.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			0: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("strconv"),
					Sel: ast.NewIdent("Atoi"),
				},
				Args: []ast.Expr{
					0: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("c"),
							Sel: ast.NewIdent("Param"),
						},
						Args: []ast.Expr{
							0: &ast.BasicLit{
								Kind:  token.STRING,
								Value: strings.ToLower(cg.entAST.Fields[0].Name),
							},
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckGetParamRead() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.DEFINE,
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

func (cg *ControllerHTTPGenerator) genCallServiceRead() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent(strings.ToLower(cg.entAST.Name)),
			1: ast.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			0: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
						Sel: ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0])),
					},
					Sel: ast.NewIdent("Get" + cg.entAST.Name + "By" + cg.entAST.Fields[0].Name),
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
					1: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "ID"),
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckCallServiceRead() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("err"),
			Op: token.DEFINE,
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

func (cg *ControllerHTTPGenerator) genSendResponseRead() *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("c"),
				Sel: ast.NewIdent("JSON"),
			},
			Args: []ast.Expr{
				0: &ast.SelectorExpr{
					X:   ast.NewIdent("http"),
					Sel: ast.NewIdent("StatusOK"),
				},
				1: ast.NewIdent(strings.ToLower(cg.entAST.Name)),
			},
		},
	}
}
