package repo_gen

// в этом файле генерация репозитрия, но мб потом вынесу генерацию в другие файли для каждого метода

import (
	"genos/internal/domain/dsl"
	"go/ast"
	"go/token"
	"strings"
)

type RepositoryGenerator struct {
	entAST            *dsl.Ent
	moduleName        string
	fullPathToPackage string
}

func NewRepositoryGenerator(moduleName string) *RepositoryGenerator {
	return &RepositoryGenerator{
		moduleName:        moduleName,
		fullPathToPackage: "internal/usecase/repo/",
	}
}

func (rg *RepositoryGenerator) GetMapRepoAST(dslAST *dsl.AST) map[string]*ast.File {
	resultMap := make(map[string]*ast.File, len(dslAST.Entities))
	for _, entityAST := range dslAST.Entities {
		resultMap[rg.getFullPathToFile(strings.ToLower(entityAST.Name))] = rg.genRepoAST(&entityAST)
	}
	return resultMap
}

func (rg *RepositoryGenerator) genRepoAST(entAST *dsl.Ent) *ast.File {
	rg.entAST = entAST
	resultAST := rg.repositoryBaseAST()
	for _, action := range entAST.Actions {
		switch action {
		case "create":
			resultAST.Decls = append(resultAST.Decls, rg.genCreateRepoAST())
		case "read":
			resultAST.Decls = append(resultAST.Decls, rg.genReadRepoAST())
		case "update":
			resultAST.Decls = append(resultAST.Decls, rg.genUpdateRepoAST())
		case "delete":
			resultAST.Decls = append(resultAST.Decls, rg.genDeleteRepoAST())
		case "list":
			resultAST.Decls = append(resultAST.Decls, rg.genListRepoAST())
		}
	}
	return resultAST
}

// генерация базовго кода для файла репозитория (без самих ручек)
func (rg *RepositoryGenerator) repositoryBaseAST() *ast.File {
	moduleName := "lol" // FIXME
	entityName := rg.entAST.Name

	return &ast.File{
		Name: ast.NewIdent("repo"),
		Decls: []ast.Decl{
			0: &ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					0: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"context\"",
						},
					},
					1: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"fmt\"",
						},
					},
					2: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/internal/entity\"",
						},
					},
					3: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/internal/usecase\"",
						},
					},
					4: &ast.ImportSpec{
						Path: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + moduleName + "/pkg/postgres\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent(entityName + "Repo"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											ast.NewIdent("pg"),
										},
										Type: &ast.StarExpr{
											X: &ast.SelectorExpr{
												X:   ast.NewIdent("postgres"),
												Sel: ast.NewIdent("Postgres"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("New" + entityName + "Repo"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("pg"),
								},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("postgres"),
										Sel: ast.NewIdent("Postgres"),
									},
								},
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: ast.NewIdent(entityName + "Repo"),
								},
							},
						},
					},
				},
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ReturnStmt{
							Results: []ast.Expr{
								&ast.UnaryExpr{
									Op: token.AND,
									X: &ast.CompositeLit{
										Type: ast.NewIdent(entityName + "Repo"),
										Elts: []ast.Expr{
											&ast.KeyValueExpr{
												Key:   ast.NewIdent("pg"),
												Value: ast.NewIdent("pg"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			3: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{
							ast.NewIdent("_"),
						},
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("usecase"),
							Sel: ast.NewIdent(entityName + "Repo"),
						},
						Values: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.ParenExpr{
									X: &ast.StarExpr{
										X: ast.NewIdent(entityName + "Repo"),
									},
								},
								Args: []ast.Expr{
									ast.NewIdent("nil"),
								},
							},
						},
					},
				},
			},
		},
	}
}

func (rg *RepositoryGenerator) getFullPathToFile(nameEntity string) string {
	return strings.Join([]string{rg.fullPathToPackage, nameEntity, "_pg.go"}, "")
}
