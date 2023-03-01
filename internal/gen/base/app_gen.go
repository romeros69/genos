package base

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func createAppAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("app"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
				},
			},
			&ast.FuncDecl{
				Name: ast.NewIdent("Run"),
				Type: &ast.FuncType{},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("fmt"),
									Sel: ast.NewIdent("Println"),
								},
								Args: []ast.Expr{
									&ast.BasicLit{
										Kind:  token.STRING,
										Value: "\"Run complete!\"",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// GenApp - Генерация App.go
func GenApp() error {
	f := createAppAST()
	fset := token.NewFileSet()
	file, err := os.Create("internal/app/app.go")
	if err != nil {
		return fmt.Errorf("error creating app.go file: %w", err)
	}
	defer file.Close()
	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in generate app.go: %w", err)
	}
	return nil
}
