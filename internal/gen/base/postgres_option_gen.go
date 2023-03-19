package base

import (
	"fmt"
	"genos/internal/util"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type PostgresOptionGenerator struct {
	file           *os.File
	moduleName     string
	fullPathToFile string
	fileAST        *ast.File
}

func NewPostgresOptionGenerator(moduleName string) *PostgresOptionGenerator {
	return &PostgresOptionGenerator{
		moduleName:     moduleName,
		fullPathToFile: "pkg/postgres/options.go",
	}
}

var _ Generator = (*PostgresOptionGenerator)(nil)

func (po *PostgresOptionGenerator) GenerateCode() error {
	err := po.preGen()
	if err != nil {
		return err
	}
	po.fileAST = createPostgresOptionsAST()
	fset := token.NewFileSet()

	err = printer.Fprint(po.file, fset, po.fileAST)
	if err != nil {
		return fmt.Errorf("error in generate %s: %w", po.file.Name(), err)
	}
	err = po.afterGen()
	if err != nil {
		return err
	}
	return nil
}

func createPostgresOptionsAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("postgres"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
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
											X: ast.NewIdent("Postgres"),
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("MaxPoolSize"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("size"),
								},
								Type: ast.NewIdent("int32"),
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
														ast.NewIdent("c"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Postgres"),
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
														X:   ast.NewIdent("c"),
														Sel: ast.NewIdent("maxPoolSize"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("size"),
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
			3: &ast.FuncDecl{
				Name: ast.NewIdent("ConnAttempts"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("attempts"),
								},
								Type: ast.NewIdent("int"),
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
														ast.NewIdent("c"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Postgres"),
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
														X:   ast.NewIdent("c"),
														Sel: ast.NewIdent("connAttempts"),
													},
												},
												Tok: token.ASSIGN,
												Rhs: []ast.Expr{
													ast.NewIdent("attempts"),
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
				Name: ast.NewIdent("ConnTimeout"),
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
														ast.NewIdent("c"),
													},
													Type: &ast.StarExpr{
														X: ast.NewIdent("Postgres"),
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
														X:   ast.NewIdent("c"),
														Sel: ast.NewIdent("connTimeout"),
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

func (po *PostgresOptionGenerator) preGen() error {
	var err error
	po.file, err = os.Create(po.fullPathToFile)
	if err != nil {
		return err
	}
	return nil
}

func (po *PostgresOptionGenerator) afterGen() error {
	// close file
	err := po.file.Close()
	if err != nil {
		return fmt.Errorf("error in closing file: %w", err)
	}

	// download dependency
	err = util.DownloadDependency(po.fileAST)
	if err != nil {
		return fmt.Errorf("error in download dependency: %w", err)
	}

	// format code
	err = util.FormatCode(po.fullPathToFile)
	if err != nil {
		return err
	}
	return nil
}
