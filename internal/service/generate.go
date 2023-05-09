package service

import (
	"fmt"
	"genos/internal/domain/dsl"
	"genos/internal/domain/principal"
	"genos/internal/domain/principal/repo_gen"
	"genos/internal/domain/principal/usecase_gen"
)

type GenerateUC struct {
	fs  FileSourceWorker
	cli CliCommandContract
}

func NewGenerateUC(fs FileSourceWorker, cli CliCommandContract) *GenerateUC {
	return &GenerateUC{fs: fs, cli: cli}
}

var _ GenerateContract = (*GenerateUC)(nil)

// GenerateDo Команда генерации crudl
func (g *GenerateUC) GenerateDo(nameModule, fullPathToDSLFile string) error {
	dslAST, err := g.parseDSL(fullPathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in GenerateDo: %w", err)
	}
	err = g.generateEntities(&dslAST)
	if err != nil {
		return fmt.Errorf("error in GenerateDo: %w", err)
	}
	err = g.generateRepository(nameModule, &dslAST)
	if err != nil {
		return fmt.Errorf("error in GenerateDo: %w", err)
	}
	err = g.generateUseCase(nameModule, &dslAST)
	if err != nil {
		return fmt.Errorf("error in GenerateDo: %w", err)
	}
	return nil
}

// функция парсинга dsl в дерево разбора
func (g *GenerateUC) parseDSL(fullPathToDSLFile string) (dsl.AST, error) {
	// Открыть файл
	file, err := g.fs.OpenFile(fullPathToDSLFile)
	if err != nil {
		return dsl.AST{}, fmt.Errorf("error in parseDSL: %w", err)
	}

	// перевести его в массив байт
	dslBytes, err := g.fs.ReadFile(file)
	if err != nil {
		return dsl.AST{}, fmt.Errorf("error in parseDSL: %w", err)
	}

	// закрыть файл
	err = g.fs.CloseFile(file)
	if err != nil {
		return dsl.AST{}, fmt.Errorf("error in parseDSL: %w", err)
	}

	// распарсить его в ast dsl
	lexer := dsl.NewLex(dslBytes)
	dsl.Parse(lexer)

	return lexer.GetResult(), nil
}

// generateEntities генерирует слой entity
func (g *GenerateUC) generateEntities(dslAST *dsl.AST) error {
	entityGenerator := principal.NewEntityGenerator()
	entityMapAST := entityGenerator.GetMapAST(dslAST)
	for path, entityAST := range entityMapAST {
		file, err := g.fs.CreateFile(path)
		if err != nil {
			return fmt.Errorf("error in generateEntities: %w", err)
		}
		err = g.fs.WriteAST(file, entityAST)
		if err != nil {
			return fmt.Errorf("error in generateEntities: %w", err)
		}
		err = g.fs.CloseFile(file)
		if err != nil {
			return fmt.Errorf("error in generateEntities: %w", err)
		}
		err = g.cli.Format(path)
		if err != nil {
			return fmt.Errorf("error in generateEntities: %w", err)
		}
	}
	return nil
}

// generateRepository генерирует слой репозитория
func (g *GenerateUC) generateRepository(nameModule string, dslAST *dsl.AST) error {
	repoGenerator := repo_gen.NewRepositoryGenerator(nameModule)
	repoMapAST := repoGenerator.GetMapRepoAST(dslAST)
	for path, repoAST := range repoMapAST {
		file, err := g.fs.CreateFile(path)
		if err != nil {
			return fmt.Errorf("error in generateRepository: %w", err)
		}
		err = g.fs.WriteAST(file, repoAST)
		if err != nil {
			return fmt.Errorf("error in generateRepository: %w", err)
		}
		err = g.fs.CloseFile(file)
		if err != nil {
			return fmt.Errorf("error in generateRepository: %w", err)
		}
		err = g.cli.Format(path)
		if err != nil {
			return fmt.Errorf("error in generateRepository: %w", err)
		}
	}
	return nil
}

func (g *GenerateUC) generateUseCase(nameModule string, dslAST *dsl.AST) error {
	ucGeneratot := usecase_gen.NewUseCaseGenerator(nameModule)
	ucMapAST := ucGeneratot.GetMapUseCaseAST(dslAST)
	for path, ucAST := range ucMapAST {
		file, err := g.fs.CreateFile(path)
		if err != nil {
			return fmt.Errorf("error in generateUseCase: %w", err)
		}
		err = g.fs.WriteAST(file, ucAST)
		if err != nil {
			return fmt.Errorf("error in generateUseCase: %w", err)
		}
		err = g.fs.CloseFile(file)
		if err != nil {
			return fmt.Errorf("error in generateUseCase: %w", err)
		}
		err = g.cli.Format(path)
		if err != nil {
			return fmt.Errorf("error in generateUseCase: %w", err)
		}
	}
	return nil
}
