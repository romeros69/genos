package command

import (
	"bufio"
	"fmt"
	"genos/internal/service"
	"os"
	"strings"
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
	err = g.checkFormatDSLFile(pathToDSLFile)
	if err != nil {
		return err
	}
	err = g.checkCurrentWorkDir()
	if err != nil {
		return fmt.Errorf("error in command Generate Do(): %w", err)
	}
	err = g.checkExistsGoSum()
	if err != nil {
		return err
	}
	nameModule, err := g.getModuleName()
	if err != nil {
		return err
	}
	err = g.genUC.GenerateDo(nameModule, pathToDSLFile)
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

func (g *Generate) checkFormatDSLFile(pathToDSLFile string) error {
	slicePath := strings.Split(pathToDSLFile, "/")
	fileName := slicePath[len(slicePath)-1]
	flag := strings.Contains(fileName, ".gek")
	if !flag {
		return fmt.Errorf("invalid format dsl file. dsl file must have .gek format")
	}
	return nil
}

func (g *Generate) checkCurrentWorkDir() error {
	_, err := os.Stat("internal/entity")
	if err != nil {
		return fmt.Errorf("error on get stat checkCurrentWorkDir: %w", err)
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("you are not at the root of the project")
	}
	return nil
}

func (g *Generate) checkExistsGoSum() error {
	_, err := os.Stat("go.mod")
	if err != nil {
		return fmt.Errorf("error in get stat file go.mod: %w", err)
	}
	if os.IsNotExist(err) {
		return fmt.Errorf("go module not found")
	}
	return nil
}

func (g *Generate) getModuleName() (string, error) {
	file, err := os.Open("go.mod") // открытие файла
	if err != nil {                // обработка ошибок
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error in closing go.mod")
		}
	}(file)

	scanner := bufio.NewScanner(file) // создание сканера

	if scanner.Scan() { // чтение первой строки
		return strings.Split(scanner.Text(), " ")[1], nil // вывод строки в консоль
	}
	return "", fmt.Errorf("not found module")
}

func (g *Generate) GetNames() []string {
	return []string{"-g", "--generate"}
}
