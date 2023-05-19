package controller_gen

import (
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genListControllerAST() *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{
						0: ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
					},
					Type: &ast.StarExpr{
						X: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Routes"),
					},
				},
			},
		},
		Name: ast.NewIdent("get" + cg.entAST.Name + "List"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{
							0: ast.NewIdent("c"),
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
				0: cg.genCallServiceList(),
				1: cg.genCheckCallServiceList(),
				2: cg.genDefineResponseList(),
				3: cg.genCheckNilResponseList(),
				4: cg.genSendResponseList(),
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCallServiceList() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			0: ast.NewIdent("list" + cg.entAST.Name),
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
					Sel: ast.NewIdent("Get" + cg.entAST.Name + "List"),
				},
				Args: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X: &ast.SelectorExpr{
								X:   ast.NewIdent("c"),
								Sel: ast.NewIdent("Request"),
							},
							Sel: ast.NewIdent("Context"),
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckCallServiceList() *ast.IfStmt {
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

func (cg *ControllerHTTPGenerator) genDefineResponseList() *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						ast.NewIdent("res"),
					},
					Type: &ast.ArrayType{
						Elt: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(cg.entAST.Name),
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckNilResponseList() *ast.IfStmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  ast.NewIdent("list" + cg.entAST.Name),
			Op: token.NEQ,
			Y:  ast.NewIdent("nil"),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("res"),
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						ast.NewIdent("list" + cg.entAST.Name),
					},
				},
			},
		},
		Else: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("res"),
					},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: ast.NewIdent("make"),
							Args: []ast.Expr{
								0: &ast.ArrayType{
									Elt: &ast.SelectorExpr{
										X:   ast.NewIdent("entity"),
										Sel: ast.NewIdent(cg.entAST.Name),
									},
								},
								1: &ast.BasicLit{
									Kind:  token.INT,
									Value: "0",
								},
							},
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genSendResponseList() *ast.ExprStmt {
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
				1: ast.NewIdent("res"),
			},
		},
	}
}
