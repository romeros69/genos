package controller_gen

import (
	"genos/internal/domain/dsl"
	"go/ast"
	"go/token"
	"strings"
)

func (cg *ControllerHTTPGenerator) genCallHandlerRouterAST(dslAST *dsl.AST, prevAST *ast.File) *ast.File {
	funcDecl, _ := prevAST.Decls[1].(*ast.FuncDecl)
	for _, entity := range dslAST.Entities {
		funcDecl.Type.Params.List = append(funcDecl.Type.Params.List, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent(strings.ToLower(entity.Name) + "Handler"),
			},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("usecase"),
				Sel: ast.NewIdent(entity.Name + "Contract"),
			},
		})
	}
	asStmt, _ := funcDecl.Body.List[0].(*ast.AssignStmt)
	asStmt.Lhs[0] = ast.NewIdent("h")
	asStmt.Tok = token.DEFINE
	constructors, _ := funcDecl.Body.List[1].(*ast.BlockStmt)
	for _, entity := range dslAST.Entities {
		constructors.List = append(constructors.List, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: ast.NewIdent("new" + entity.Name + "Routes"),
				Args: []ast.Expr{
					0: ast.NewIdent("h"),
					1: ast.NewIdent(strings.ToLower(entity.Name) + "Handler"),
				},
			},
		})
	}
	return prevAST
}
