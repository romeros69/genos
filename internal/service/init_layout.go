package service

import (
	"fmt"
	"genos/internal/domain"
	"genos/internal/util"
	"os/exec"
)

type GenerateBase struct {
	fw FileSourceWorker
}

func NewGenerateSource(fw FileSourceWorker) *GenerateBase {
	return &GenerateBase{fw: fw}
}

var _ GenerateBaseContract = (*GenerateBase)(nil)

func (gs *GenerateBase) initBaseGenerators(moduleName string) []domain.BaseGenerator {
	return []domain.BaseGenerator{
		0: domain.NewPostgresOptionGenerator(),
		1: domain.NewHttpOptionsGenerator(),
		2: domain.NewHttpServerGenerator(),
		3: domain.NewPostgresGenerator(moduleName),
		4: domain.NewConfigGenerator(),
		5: domain.NewAppGenerator(moduleName),
		6: domain.NewMainGenerator(moduleName),
	}
}

// Генерация базового кода - все таки это только часть выполнения определенной команды
func (gs *GenerateBase) generateBaseCode(moduleName string) error {
	baseGenerators := gs.initBaseGenerators(moduleName)
	for _, v := range baseGenerators {
		file, err := gs.fw.CreateFile(v.FullPathToFile())
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
		newAST := v.GenAST()
		err = gs.fw.WriteAST(file, newAST)
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
		err = gs.fw.CloseFile(file)
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
		err = util.DownloadDependency(newAST)
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
		err = util.FormatCode(v.FullPathToFile())
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
	}
	return nil
}

func (gs *GenerateBase) createGoModule(nameProject string) error {
	cmd := exec.Command("go", "mod", "init", nameProject)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error in execute go mod init: %w", err)
	}
	return nil
}