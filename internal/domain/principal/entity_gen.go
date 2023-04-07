package principal

import (
	"genos/internal/domain/dsl"
	"genos/internal/util"
	"go/ast"
	"go/token"
	"strings"
)

type EntityGenerator struct {
	fullPathToPackage string
}

func NewEntityGenerator() *EntityGenerator {
	return &EntityGenerator{
		fullPathToPackage: "internal/entity/",
	}
}

func (eg *EntityGenerator) GetMapAST(astDSL *dsl.AST) map[string]*ast.File {
	resultMap := make(map[string]*ast.File, len(astDSL.Entities))
	for _, v := range astDSL.Entities {
		resultMap[eg.getFullPathToFile(strings.ToLower(v.Name), eg.fullPathToPackage)] = eg.genAST(v)
	}
	return resultMap
}

// точка входа в генерацию сущности
func (eg *EntityGenerator) genAST(ent dsl.Ent) *ast.File {
	return &ast.File{
		Name: ast.NewIdent("entity"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.TYPE,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent(ent.Name),
						Type: &ast.StructType{Fields: eg.getFieldsList(ent.Fields)},
					},
				},
			},
		},
	}
}

func (eg *EntityGenerator) getFieldsList(fieldsDSL []dsl.Field) *ast.FieldList {
	listField := make([]*ast.Field, len(fieldsDSL))
	for i, v := range fieldsDSL {
		listField[i] = &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent(v.Name),
			},
			Type: ast.NewIdent(util.TypesMap[v.TokType]),
		}
	}
	return &ast.FieldList{List: listField}
}

func (eg *EntityGenerator) getFullPathToFile(nameEntity, fullPathToPackage string) string {
	return fullPathToPackage + nameEntity + ".go"
}
