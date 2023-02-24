package generate

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func GenMain(nameModule string) error {
	f := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + nameModule + "/internal/" + "app\"",
						},
					},
				},
			},
			&ast.FuncDecl{
				Name: ast.NewIdent("main"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						Opening: 10,
						Closing: 11,
					},
				},
				Body: &ast.BlockStmt{
					//Lbrace: 13,
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("app"),
									Sel: ast.NewIdent("Run"),
								},
								//Args: []ast.Expr{
								//	&ast.BasicLit{
								//		Kind:  token.STRING,
								//		Value: "\"hello, it is genos\"",
								//	},
								//},
							},
						},
					},
				},
			},
		},
	}
	fset := token.NewFileSet()

	file, err := os.Create("cmd/main/main.go")
	if err != nil {
		return fmt.Errorf("error in creating main.go file: %w", err)
	}
	defer file.Close()

	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in genereate main: %w", err)
	}
	return nil
}

func GenApp() error {
	f := &ast.File{
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
	fset := token.NewFileSet()
	file, err := os.Create("internal/app/app.go")
	if err != nil {
		return fmt.Errorf("error creating app.go file: %w", err)
	}
	defer file.Close()
	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in generate app.go")
	}
	return nil
}
