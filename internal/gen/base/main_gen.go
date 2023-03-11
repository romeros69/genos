package base

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

func createMainAST(moduleName string) *ast.File {
	return &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"log\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/configs\"",
						},
					},
					2: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/internal/app\"",
						},
					},
				},
			},
			1: &ast.FuncDecl{
				Name: ast.NewIdent("main"),
				Type: &ast.FuncType{
					Params: nil,
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						0: &ast.AssignStmt{
							Lhs: []ast.Expr{
								0: ast.NewIdent("cfg"),
								1: ast.NewIdent("err"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("configs"),
										Sel: ast.NewIdent("NewConfig"),
									},
								},
							},
						},
						1: &ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X:  ast.NewIdent("err"),
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun: &ast.SelectorExpr{
												X:   ast.NewIdent("log"),
												Sel: ast.NewIdent("Fatalf"),
											},
											Args: []ast.Expr{
												0: &ast.BasicLit{
													Kind:  token.STRING,
													Value: "\"Error in parse config: %s\\n\"",
												},
												1: ast.NewIdent("err"),
											},
										},
									},
								},
							},
						},
						2: &ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("app"),
									Sel: ast.NewIdent("Run"),
								},
								Args: []ast.Expr{
									ast.NewIdent("cfg"),
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
func GenMain(moduleName string) error {
	f := createMainAST(moduleName)
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
