package service

import (
	"fmt"
	"genos/internal/domain"
	"genos/internal/util"
)

type InitLayoutUC struct {
	fw  FileSourceWorker
	cli CliCommandContract
	fwc FolderContract
}

func NewInitLayout(fw FileSourceWorker, cli CliCommandContract, fwc FolderContract) *InitLayoutUC {
	return &InitLayoutUC{fw: fw, cli: cli, fwc: fwc}
}

var _ InitLayoutContract = (*InitLayoutUC)(nil)

func (gs *InitLayoutUC) InitLayoutDo(nameProject, path string) error {
	err := gs.fwc.CreateFolder(nameProject, path)
	if err != nil {
		return fmt.Errorf("error InitLayutDo: %w", err)
	}
	err = gs.cli.CreateGoModule(nameProject)
	if err != nil {
		return fmt.Errorf("error InitLayutDo: %w", err)
	}
	err = gs.generateBaseCode(nameProject)
	if err != nil {
		return fmt.Errorf("error InitLayutDo: %w", err)
	}
	return nil
}

func (gs *InitLayoutUC) initBaseGenerators(moduleName string) []domain.BaseGenerator {
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
func (gs *InitLayoutUC) generateBaseCode(moduleName string) error {
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
		err = gs.cli.Format(v.FullPathToFile())
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
	}
	return nil
}
