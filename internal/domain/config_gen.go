package domain

import (
	"go/ast"
	"go/token"
)

type ConfigGenerator struct {
	fullPathToFile string
}

var _ BaseGenerator = (*ConfigGenerator)(nil)

func NewConfigGenerator() *ConfigGenerator {
	return &ConfigGenerator{
		fullPathToFile: "configs/configs.go",
	}
}

func (cg *ConfigGenerator) FullPathToFile() string {
	return cg.fullPathToFile
}

func (cg *ConfigGenerator) GenAST() *ast.File {
	return &ast.File{
		Name: ast.NewIdent("configs"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"github.com/caarlos0/env/v6\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("Config"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									0: {
										Names: []*ast.Ident{
											ast.NewIdent("AppPort"),
										},
										Type: ast.NewIdent("string"),
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`env:\"APP_PORT\" envDefault:\"9000\"`",
										},
									},
									1: {
										Names: []*ast.Ident{
											ast.NewIdent("PostgresHost"),
										},
										Type: ast.NewIdent("string"),
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`env:\"POSTGRES_HOST\" envDefault:\"localhost\"`",
										},
									},
									2: {
										Names: []*ast.Ident{
											ast.NewIdent("PostgresPort"),
										},
										Type: ast.NewIdent("string"),
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`env:\"POSTGRES_PORT\" envDefault:\"5432\"`",
										},
									},
									3: {
										Names: []*ast.Ident{
											ast.NewIdent("PostgresUser"),
										},
										Type: ast.NewIdent("string"),
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`env:\"POSTGRES_USER\" envDefault:\"postgres\"`",
										},
									},
									4: {
										Names: []*ast.Ident{
											ast.NewIdent("PostgresPassword"),
										},
										Type: ast.NewIdent("string"),
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`env:\"POSTGRES_PASSWORD\" envDefault:\"postgres\"`",
										},
									},
									5: {
										Names: []*ast.Ident{
											ast.NewIdent("PostgresDBName"),
										},
										Type: ast.NewIdent("string"),
										Tag: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "`env:\"POSTGRES_DB_NAME\" envDefault:\"postgres\"`",
										},
									},
									6: {
										Names: []*ast.Ident{
											ast.NewIdent("PostgresURL"),
										},
										Type: ast.NewIdent("string"),
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("NewConfig"),
				Type: &ast.FuncType{
					Params: nil,
					Results: &ast.FieldList{
						List: []*ast.Field{
							0: {
								Type: &ast.StarExpr{
									X: ast.NewIdent("Config"),
								},
							},
							1: {
								Type: ast.NewIdent("error"),
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						0: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("cfg"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: ast.NewIdent("new"),
									Args: []ast.Expr{
										ast.NewIdent("Config"),
									},
								},
							},
						},
						1: &ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("err"),
							},
							Tok: token.DEFINE,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("env"),
										Sel: ast.NewIdent("Parse"),
									},
									Args: []ast.Expr{
										ast.NewIdent("cfg"),
									},
								},
							},
						},
						2: &ast.IfStmt{
							Cond: &ast.BinaryExpr{
								X:  ast.NewIdent("err"),
								Op: token.NEQ,
								Y:  ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								List: []ast.Stmt{
									&ast.ReturnStmt{
										Results: []ast.Expr{
											0: ast.NewIdent("nil"),
											1: ast.NewIdent("err"),
										},
									},
								},
							},
						},
						3: &ast.AssignStmt{
							Lhs: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("cfg"),
									Sel: ast.NewIdent("PostgresURL"),
								},
							},
							Tok: token.ASSIGN,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("fmt"),
										Sel: ast.NewIdent("Sprintf"),
									},
									Args: []ast.Expr{
										0: &ast.BasicLit{
											Kind:  token.STRING,
											Value: "\"postgresql://%s:%s@%s:%s/%s\"",
										},
										1: &ast.SelectorExpr{
											X:   ast.NewIdent("cfg"),
											Sel: ast.NewIdent("PostgresUser"),
										},
										2: &ast.SelectorExpr{
											X:   ast.NewIdent("cfg"),
											Sel: ast.NewIdent("PostgresPassword"),
										},
										3: &ast.SelectorExpr{
											X:   ast.NewIdent("cfg"),
											Sel: ast.NewIdent("PostgresHost"),
										},
										4: &ast.SelectorExpr{
											X:   ast.NewIdent("cfg"),
											Sel: ast.NewIdent("PostgresPort"),
										},
										5: &ast.SelectorExpr{
											X:   ast.NewIdent("cfg"),
											Sel: ast.NewIdent("PostgresDBName"),
										},
									},
								},
							},
						},
						4: &ast.ReturnStmt{
							Results: []ast.Expr{
								0: ast.NewIdent("cfg"),
								1: ast.NewIdent("nil"),
							},
						},
					},
				},
			},
		},
	}
}
