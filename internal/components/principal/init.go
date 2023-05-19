package principal

import (
	"genos/internal/components/dsl"
	"go/ast"
	"go/token"
	"strings"
)

func GetUpdateAppInit(moduleName string, dslAST *dsl.AST, prevAST *ast.File) *ast.File {
	funcDecl, _ := prevAST.Decls[1].(*ast.FuncDecl)
	oldStmtLeft := funcDecl.Body.List[:2]
	oldStmtMiddle := funcDecl.Body.List[2:4]
	oldStmtRight := funcDecl.Body.List[4:]
	newResStmt := make([]ast.Stmt, 0)
	for _, stmt := range oldStmtLeft {
		newResStmt = append(newResStmt, stmt)
	}
	for _, entity := range dslAST.Entities {
		newResStmt = append(newResStmt, getAssignInit(entity))
	}
	for _, stmt := range oldStmtMiddle {
		newResStmt = append(newResStmt, stmt)
	}
	newResStmt = append(newResStmt, initRouter(dslAST))
	for _, stmt := range oldStmtRight {
		newResStmt = append(newResStmt, stmt)
	}
	pgStmt := newResStmt[0].(*ast.AssignStmt)
	pgStmt.Lhs[0].(*ast.Ident).Name = "pg"
	funcDecl.Body.List = newResStmt
	addImportsInit(moduleName, prevAST)
	return prevAST
}

func getAssignInit(entity dsl.Ent) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(strings.ToLower(entity.Name) + "UseCase")},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("usecase"),
					Sel: ast.NewIdent("New" + entity.Name + "UseCase"),
				},
				Args: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("repo"),
							Sel: ast.NewIdent("New" + entity.Name + "Repo"),
						},
						Args: []ast.Expr{
							ast.NewIdent("pg"),
						},
					},
				},
			},
		},
	}
}

func addImportsInit(moduleName string, fileAST *ast.File) {
	fileAST.Decls[0].(*ast.GenDecl).Specs = append(fileAST.Decls[0].(*ast.GenDecl).Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"" + moduleName + "/internal/usecase\"",
		},
	})
	fileAST.Decls[0].(*ast.GenDecl).Specs = append(fileAST.Decls[0].(*ast.GenDecl).Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"" + moduleName + "/internal/usecase/repo\"",
		},
	})
	fileAST.Decls[0].(*ast.GenDecl).Specs = append(fileAST.Decls[0].(*ast.GenDecl).Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"" + moduleName + "/internal/controller/http\"",
		},
	})
}

func initRouter(dslAST *dsl.AST) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   ast.NewIdent("http"),
				Sel: ast.NewIdent("NewRouter"),
			},
			Args: func() []ast.Expr {
				res := make([]ast.Expr, 0)
				res = append(res, ast.NewIdent("handler"))
				for _, entity := range dslAST.Entities {
					res = append(res, ast.NewIdent(strings.ToLower(entity.Name)+"UseCase"))
				}
				return res
			}(),
		},
	}
}
