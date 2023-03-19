package base

import (
	"fmt"
	"genos/internal/util"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type HttpOptionsGenerator struct {
	file           *os.File
	moduleName     string
	fullPathToFile string
	fileAST        *ast.File
}

func NewHttpOptionsGenerator(moduleName string) *HttpOptionsGenerator {
	return &HttpOptionsGenerator{
		moduleName:     moduleName,
		fullPathToFile: "pkg/httpserver/options.go",
	}
}

var _ Generator = (*HttpOptionsGenerator)(nil)

func (ho *HttpOptionsGenerator) GenerateCode() error {
	err := ho.preGen()
	if err != nil {
		return err
	}
	ho.fileAST = createHttpOptionsAST()
	fset := token.NewFileSet()

	err = printer.Fprint(ho.file, fset, ho.fileAST)
	if err != nil {
		return fmt.Errorf("error in generate %s: %w", ho.file.Name(), err)
	}
	err = ho.afterGen()
	if err != nil {
		return err
	}
	return nil
}

func createHttpOptionsAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("httpserver"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"net\"",
						},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"time\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Option"),
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{
										Type: &ast.StarExpr{
											X: ast.NewIdent("Server"),
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("Port"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("port"),
								},
								Type: ast.NewIdent("string"),
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("Addr"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													&ast.CallExpr{
														Fun: &ast.SelectorExpr{
															X:   ast.NewIdent("net"),
															Sel: ast.NewIdent("JoinHostPort"),
														},
														Args: []ast.Expr{
															0: &ast.BasicLit{
																Kind:  token.STRING,
																Value: "\"\"",
															},
															1: ast.NewIdent("port"),
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			// 3
			3: &ast.FuncDecl{
				Name: ast.NewIdent("ReadTimeout"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("timeout"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Duration"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("ReadTimeout"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("timeout"),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			4: &ast.FuncDecl{
				Name: ast.NewIdent("WriteTimeout"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("timeout"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Duration"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X: &ast.SelectorExpr{
															X:   ast.NewIdent("s"),
															Sel: ast.NewIdent("server"),
														},
														Sel: ast.NewIdent("WriteTimeout"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("timeout"),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			5: &ast.FuncDecl{
				Name: ast.NewIdent("ShutdownTimeout"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("timeout"),
								},
								Type: &ast.SelectorExpr{
									X:   ast.NewIdent("time"),
									Sel: ast.NewIdent("Duration"),
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: ast.NewIdent("Option"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.FuncLit{
									Type: &ast.FuncType{
										Params: &ast.FieldList{
											List: []*ast.Field{
												{
													Names: []*ast.Ident{
														ast.NewIdent("s"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Server"),
													},
												},
											},
										},
									},
									Body: &ast.BlockStmt{
										List: []ast.Stmt{
											&ast.AssignStmt{
												Lhs: []ast.Expr{
													&ast.SelectorExpr{
														X:   ast.NewIdent("s"),
														Sel: ast.NewIdent("shutdownTimeout"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("timeout"),
												},
											},
										},
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

func (ho *HttpOptionsGenerator) preGen() error {
	var err error
	ho.file, err = os.Create(ho.fullPathToFile)
	if err != nil {
		return err
	}
	return nil
}

func (ho *HttpOptionsGenerator) afterGen() error {
	// close file
	err := ho.file.Close()
	if err != nil {
		return fmt.Errorf("error in closing file: %w", err)
	}

	// download dependency
	err = util.DownloadDependency(ho.fileAST)
	if err != nil {
		return fmt.Errorf("error in download dependency: %w", err)
	}

	// format code
	err = util.FormatCode(ho.fullPathToFile)
	if err != nil {
		return err
	}
	return nil
}
