package base

import (
	"fmt"
	"genos/internal/util"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type MainGenerator struct {
	file           *os.File
	moduleName     string
	fullPathToFile string
	fileAST        *ast.File
}

func NewMainGenerator(moduleName string) *MainGenerator {
	return &MainGenerator{
		moduleName:     moduleName,
		fullPathToFile: "cmd/main/main.go",
	}
}

var _ Generator = (*MainGenerator)(nil)

func (mg *MainGenerator) GenerateCode() error {
	err := mg.preGen()
	if err != nil {
		return err
	}
	mg.fileAST = createMainAST(mg.moduleName)
	fset := token.NewFileSet()
	err = printer.Fprint(mg.file, fset, mg.fileAST)
	if err != nil {
		return fmt.Errorf("error in generate %s: %w", mg.file.Name(), err)
	}
	err = mg.afterGen()
	if err != nil {
		return err
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

func (mg *MainGenerator) preGen() error {
	var err error
	mg.file, err = os.Create(mg.fullPathToFile)
	if err != nil {
		return err
	}
	return nil
}

func (mg *MainGenerator) afterGen() error {
	// close file
	err := mg.file.Close()
	if err != nil {
		return fmt.Errorf("error in closing file: %w", err)
	}

	// download dependency
	err = util.DownloadDependency(mg.fileAST)
	if err != nil {
		return fmt.Errorf("error in download dependency: %w", err)
	}

	// format code
	err = util.FormatCode(mg.fullPathToFile)
	if err != nil {
		return err
	}
	return nil
}
