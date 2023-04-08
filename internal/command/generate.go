package command

import (
	"fmt"
	"genos/internal/service"
	"os"
)

type Generate struct {
	genUC service.GenerateContract
}

func NewGenerate(genUC service.GenerateContract) *Generate {
	return &Generate{genUC: genUC}
}

// Do TODO добавить проверки пути + установка рабочей директории
func (g *Generate) Do(pathToDSLFile string) error {
	err := g.checkExistFile(pathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in command Generate Do(): %w", err)
	}
	err = g.genUC.GenerateDo(pathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in command Generate Do(): %w", err)
	}
	return nil
}

func (g *Generate) checkExistFile(pathToDSLFile string) error {
	// проверяем что файл существует
	_, err := os.Stat(pathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in get stat file %s: %w", pathToDSLFile, err)
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist %s: %w", pathToDSLFile, err)
	}
	return nil
}

func (g *Generate) GetNames() []string {
	return []string{"-g", "--generate"}
}
