package base

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type MainGenerator struct {
	moduleName string
	fileName   string
	file       *os.File
	fileAST    *ast.File
}

func NewMainGenerator(moduleName string) *MainGenerator {
	return &MainGenerator{
		moduleName: moduleName,
		fileName:   "cmd/main/main.go",
	}
}

var _ Generator = (*MainGenerator)(nil)

func (mg *MainGenerator) GenerateCode() error {
	mg.fileAST = createMainAST(mg.moduleName)
	fset := token.NewFileSet()
	err := printer.Fprint(mg.file, fset, mg.fileAST)
	if err != nil {
		return fmt.Errorf("error in genereate main: %w", err)
	}
	return nil
}

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
