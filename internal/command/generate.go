package command

import (
	"fmt"
	"genos/internal/service"
	"os"
)

type Generate struct {
	genUC             service.GenerateContract
	fullPathToDSLFile string
}

func NewGenerate(genUC service.GenerateContract) *Generate {
	return &Generate{genUC: genUC}
}

// Do TODO добавить проверки пути + установка рабочей директории
func (g *Generate) Do() error {
	err := g.readPathToDSL()
	if err != nil {
		return fmt.Errorf("error in command Generate Do(): %w", err)
	}
	err = g.genUC.GenerateDo(g.fullPathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in command Generate Do(): %w", err)
	}
	return nil
}

func (g *Generate) readPathToDSL() error {
	// читаем путь к файлу dsl
	fmt.Printf("enter full path to dsl file:\n")
	_, err := fmt.Scanln(&g.fullPathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in reading path to dsl file: %w", err)
	}
	// проверяем что файл существует
	_, err = os.Stat(g.fullPathToDSLFile)
	if err != nil {
		return fmt.Errorf("error in get stat file %s: %w", g.fullPathToDSLFile, err)
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("file doesn't exist %s: %w", g.fullPathToDSLFile, err)
	}
	return nil
}

func (g *Generate) GetNames() []string {
	return []string{"-g", "--generate"}
}
