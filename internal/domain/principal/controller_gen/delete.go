package controller_gen

import (
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genDeleteControllerAST() *ast.FuncDecl {
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
		Name: ast.NewIdent("delete" + cg.entAST.Name),
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
				0: cg.genGetParamQueryDelete(),
				1: cg.genCheckGetParamDelete(),
				2: cg.genCallServiceDelete(),
				3: cg.genCheckCallServiceDelete(),
				4: cg.genSendResponseDelete(),
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genGetParamQueryDelete() *ast.AssignStmt {
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

func (cg *ControllerHTTPGenerator) genCheckGetParamDelete() *ast.IfStmt {
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
							2: &ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"not found\"",
							},
						},
					},
				},
				1: &ast.ReturnStmt{},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCallServiceDelete() *ast.AssignStmt {
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
					Sel: ast.NewIdent("Delete" + cg.entAST.Name),
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
					1: &ast.CallExpr{
						Fun: ast.NewIdent(util.TypesMap[cg.entAST.Fields[0].TokType]),
						Args: []ast.Expr{
							ast.NewIdent(strings.ToLower(cg.entAST.Name) + "ID"),
						},
					},
				},
			},
		},
	}
}

func (cg *ControllerHTTPGenerator) genCheckCallServiceDelete() *ast.IfStmt {
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

func (cg *ControllerHTTPGenerator) genSendResponseDelete() *ast.ExprStmt {
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
