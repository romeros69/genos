package usecase_gen

import (
	"genos/internal/components/dsl"
	"go/ast"
	"go/token"
	"strings"
)

type UseCaseGenerator struct {
	entAST            *dsl.Ent
	moduleName        string
	fullPathToPackage string
}

func NewUseCaseGenerator(moduleName string) *UseCaseGenerator {
	return &UseCaseGenerator{
		moduleName:        moduleName,
		fullPathToPackage: "internal/usecase/",
	}
}

func (ug *UseCaseGenerator) GetMapUseCaseAST(dslAST *dsl.AST) map[string]*ast.File {
	resultMap := make(map[string]*ast.File, len(dslAST.Entities)+1)
	for _, entityAST := range dslAST.Entities {
		resultMap[ug.getFullPathToFile(strings.ToLower(entityAST.Name))] = ug.genUseCaseAST(&entityAST)
	}
	resultMap[ug.getFullPathToFile("interfaces")] = ug.genInterfaceAST(dslAST)
	return resultMap
}

func (ug *UseCaseGenerator) genUseCaseAST(entAST *dsl.Ent) *ast.File {
	ug.entAST = entAST
	resultAST := ug.useCaseBaseAST()
	for _, action := range entAST.Actions {
		switch action {
		case "create":
			resultAST.Decls = append(resultAST.Decls, ug.genCreateUCAST())
		case "read":
			resultAST.Decls = append(resultAST.Decls, ug.genReadUCAST())
		case "update":
			resultAST.Decls = append(resultAST.Decls, ug.genUpdateUCAST())
		case "delete":
			resultAST.Decls = append(resultAST.Decls, ug.genDeleteUCAST())
		case "list":
			resultAST.Decls = append(resultAST.Decls, ug.genListUCAST())
		}
	}
	return resultAST
}

func (ug *UseCaseGenerator) useCaseBaseAST() *ast.File {
	moduleName := ug.moduleName
	entityName := ug.entAST.Name
	return &ast.File{
		Name: ast.NewIdent("usecase"),
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
							Value: "\"" + moduleName + "/internal/entity\"",
						},
					},
				},
			},
			1: &ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent(entityName + "UseCase"),
						Type: &ast.StructType{
							Fields: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{
											ast.NewIdent("repo"),
										},
										Type: ast.NewIdent(entityName + "Repo"),
									},
								},
							},
						},
					},
				},
			},
			2: &ast.FuncDecl{
				Name: ast.NewIdent("New" + entityName + "UseCase"),
				Type: &ast.FuncType{
					Params: &ast.FieldList{
						List: []*ast.Field{
							{
								Names: []*ast.Ident{
									ast.NewIdent("r"),
								},
								Type: ast.NewIdent(entityName + "Repo"),
							},
						},
					},
					Results: &ast.FieldList{
						List: []*ast.Field{
							{
								Type: &ast.StarExpr{
									X: ast.NewIdent(entityName + "UseCase"),
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
										Type: ast.NewIdent(entityName + "UseCase"),
										Elts: []ast.Expr{
											&ast.KeyValueExpr{
												Key:   ast.NewIdent("repo"),
												Value: ast.NewIdent("r"),
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
						Type: ast.NewIdent(entityName + "Contract"),
						Values: []ast.Expr{
							&ast.CallExpr{
								Fun: &ast.ParenExpr{
									X: &ast.StarExpr{
										X: ast.NewIdent(entityName + "UseCase"),
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

func (ug *UseCaseGenerator) getFullPathToFile(nameEntity string) string {
	return strings.Join([]string{ug.fullPathToPackage, nameEntity, ".go"}, "")
}
