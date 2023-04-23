package usecase_gen

import (
	"fmt"
	"genos/internal/domain/dsl"
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

func (ug *UseCaseGenerator) genInterfaceAST(dslAST *dsl.AST) *ast.File {
	resultAST := ug.genBaseInterfacesAST()
	for _, entityAST := range dslAST.Entities {
		fmt.Println("1111111")
		fmt.Println(entityAST.Name)
		fmt.Println("2222222")
		resultAST.Decls = append(resultAST.Decls, ug.genInterfaceDeclByEntity(&entityAST, "Contract"))
		resultAST.Decls = append(resultAST.Decls, ug.genInterfaceDeclByEntity(&entityAST, "Repo"))
	}
	return resultAST
}

func (ug *UseCaseGenerator) genBaseInterfacesAST() *ast.File {
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
							Value: strings.Join([]string{"\"", ug.moduleName, "/internal/entity\""}, ""),
						},
					},
				},
			},
		},
	}
}

func (ug *UseCaseGenerator) genInterfaceDeclByEntity(entAST *dsl.Ent, typeInterface string) *ast.GenDecl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(strings.Join([]string{entAST.Name, typeInterface}, "")),
				Type: &ast.InterfaceType{
					Methods: &ast.FieldList{
						List: func() []*ast.Field {
							res := make([]*ast.Field, 0)
							for _, action := range entAST.Actions {
								switch action {
								case "create":
									res = append(res, ug.genCreateMethodInterface(entAST))
								case "read":
									res = append(res, ug.genReadMethodInterface(entAST))
								case "update":
									res = append(res, ug.genUpdateMethodInterface(entAST))
								case "delete":
									res = append(res, ug.genDeleteMethodInterface(entAST))
								case "list":
									res = append(res, ug.genListMethodInterface(entAST))
								}
							}
							return res
						}(),
					},
				},
			},
		},
	}
}

func (ug *UseCaseGenerator) genCreateMethodInterface(entAST *dsl.Ent) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(strings.Join([]string{"Create", entAST.Name}, "")),
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
					1: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(entAST.Name),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: ast.NewIdent(util.TypesMap[entAST.Fields[0].TokType]),
					},
					1: {
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
	}
}

func (ug *UseCaseGenerator) genReadMethodInterface(entAST *dsl.Ent) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(strings.Join([]string{"Get", entAST.Name, "By", entAST.Fields[0].Name}, "")),
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
					1: {
						Type: ast.NewIdent(util.TypesMap[entAST.Fields[0].TokType]),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(entAST.Name),
						},
					},
					1: {
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
	}
}

func (ug *UseCaseGenerator) genUpdateMethodInterface(entAST *dsl.Ent) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(strings.Join([]string{"Update", entAST.Name}, "")),
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
					1: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("entity"),
							Sel: ast.NewIdent(entAST.Name),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
	}
}

// index 3
func (ug *UseCaseGenerator) genDeleteMethodInterface(entAST *dsl.Ent) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(strings.Join([]string{"Delete", entAST.Name}, "")),
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
					1: {
						Type: ast.NewIdent(util.TypesMap[entAST.Fields[0].TokType]),
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
	}
}

func (ug *UseCaseGenerator) genListMethodInterface(entAST *dsl.Ent) *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(strings.Join([]string{"Get", entAST.Name, "List"}, "")),
		},
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("context"),
							Sel: ast.NewIdent("Context"),
						},
					},
				},
			},
			Results: &ast.FieldList{
				List: []*ast.Field{
					0: {
						Type: &ast.ArrayType{
							Elt: &ast.SelectorExpr{
								X:   ast.NewIdent("entity"),
								Sel: ast.NewIdent(entAST.Name),
							},
						},
					},
					1: {
						Type: ast.NewIdent("error"),
					},
				},
			},
		},
	}
}

func (ug *UseCaseGenerator) genInterfaceBaseAST() *ast.File {
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
							Value: "\"" + ug.moduleName + "/internal/entity\"",
						},
					},
				},
			},
		},
	}
}
