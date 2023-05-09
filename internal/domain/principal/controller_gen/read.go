package controller_gen

//
//import (
//	"go/ast"
//	"strings"
//)
//
//func (cg *ControllerHTTPGenerator) genReadControllerAST() *ast.FuncDecl {
//	return &ast.FuncDecl{
//		Recv: &ast.FieldList{
//			List: []*ast.Field{
//				{
//					Names: []*ast.Ident{
//						ast.NewIdent(string(strings.ToLower(cg.entAST.Name)[0]) + "r"),
//					},
//					Type: &ast.StarExpr{
//						X: ast.NewIdent(strings.ToLower(cg.entAST.Name) + "Routes"),
//					},
//				},
//			},
//		},
//		Name: ast.NewIdent("get" + cg.entAST.Name + "By" + cg.entAST.Fields[0].Name),
//		Type: &ast.FuncType{
//			Params: &ast.FieldList{
//				List: []*ast.Field{
//					{
//						Names: []*ast.Ident{
//							ast.NewIdent("c"),
//						},
//						Type: &ast.StarExpr{
//							X: &ast.SelectorExpr{
//								X: ast.NewIdent("gin"),
//								Sel: ast.NewIdent("Context"),
//							},
//						},
//					},
//				},
//			},
//		},
//		Body: &ast.BlockStmt{
//			List: []ast.Stmt{
//				0: ,
//				1: ,
//				2: ,
//				3: ,
//				4: ,
//			},
//		},
//	}
//}
//
//func (cg *ControllerHTTPGenerator) genGetParamQueryRead() *ast.AssignStmt {
//
//}
//
//func (cg *ControllerHTTPGenerator) genCheckGetParamRead() *ast.IfStmt {
//
//}
//
//func (cg *ControllerHTTPGenerator) genCallServiceRead() *ast.AssignStmt {
//
//}
//
//func (cg *ControllerHTTPGenerator) genCheckCallServiceRead() *ast.IfStmt {
//
//}
//
//func (cg *ControllerHTTPGenerator) genSendResponseRead() *ast.ExprStmt {
//
//}
