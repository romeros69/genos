package service

import (
	"fmt"
	"genos/internal/components/base"
	"genos/internal/components/base/docker"
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
	err = gs.generateDocker()
	if err != nil {
		return fmt.Errorf("error InitLayutDo: %w", err)
	}
	err = gs.generateDockerCompose()
	if err != nil {
		return fmt.Errorf("error InitLayutDo: %w", err)
	}
	return nil
}

func (gs *InitLayoutUC) initBaseGenerators(moduleName string) []base.BaseGenerator {
	return []base.BaseGenerator{
		0: base.NewPostgresOptionGenerator(),
		1: base.NewHttpOptionsGenerator(),
		2: base.NewHttpServerGenerator(),
		3: base.NewPostgresGenerator(moduleName),
		4: base.NewConfigGenerator(),
		5: base.NewAppGenerator(moduleName),
		6: base.NewMainGenerator(moduleName),
		7: base.NewRouterGenerator(moduleName),
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
		//if i == 0 {
		//	fmt.Printf("Start download dependency...\n")
		//}
		//
		//err = util.DownloadDependency(newAST)
		//if err != nil {
		//	return fmt.Errorf("error in GenerateBaseCode: %w", err)
		//}
		//if i == len(baseGenerators)-1 {
		//	fmt.Printf("Complete!\n")
		//}
		err = gs.cli.Format(v.FullPathToFile())
		if err != nil {
			return fmt.Errorf("error in GenerateBaseCode: %w", err)
		}
	}
	return nil
}

func (gs *InitLayoutUC) generateDocker() error {
	file, err := gs.fw.CreateFile("Dockerfile")
	if err != nil {
		return fmt.Errorf("error in generateDocker(): %w", err)
	}
	err = gs.fw.WriteByteSliceFile(file, docker.GetDockerFile())
	err = gs.fw.CloseFile(file)
	if err != nil {
		return fmt.Errorf("error in generateDocker(): %w", err)
	}
	return nil
}

func (gs *InitLayoutUC) generateDockerCompose() error {
	file, err := gs.fw.CreateFile("docker-compose.yml")
	if err != nil {
		return fmt.Errorf("error in generateDockerCompose(): %w", err)
	}
	err = gs.fw.WriteByteSliceFile(file, docker.GetDockerComposeFile())
	err = gs.fw.CloseFile(file)
	if err != nil {
		return fmt.Errorf("error in generateDockerCompose(): %w", err)
	}
	return nil
}
