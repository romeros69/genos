package base

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func createMainAST(nameModule string) *ast.File {
	return &ast.File{
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
							},
						},
					},
				},
			},
		},
	}
}

// GenMain - генерация main.go
func GenMain(nameModule string) error {
	f := createMainAST(nameModule)
	fset := token.NewFileSet()

	file, err := os.Create("cmd/main/main.go")
	if err != nil {
		return fmt.Errorf("error in creating main.go file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error in closing file %s", file.Name())
		}
	}(file)

	err = printer.Fprint(file, fset, f)
	if err != nil {
		return fmt.Errorf("error in genereate main: %w", err)
	}
	return nil
}
