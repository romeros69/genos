package layout

import (
	"fmt"
	"genos/internal/commands"
	"io/fs"
	"os"
	"os/exec"
	"unicode/utf8"
)

// Команда для создания скилета проекта (genos -g --generate)

const perm fs.FileMode = 0777

type Project struct {
	Path string
}

func createGoModule(moduleName string) error {
	cmd := exec.Command("go", "mod", "init", moduleName)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error in execute go mod init: %w", err)
	}
	return nil
}

func createMain() (*os.File, error) {
	// mkdir cmd/main
	err := os.MkdirAll("cmd/main", perm)
	if err != nil {
		return nil, fmt.Errorf("error in creating /cmd/main directory: %w", err)
	}

	// create mian.go
	f, err := os.Create("cmd/main/main.go")
	if err != nil {
		return nil, fmt.Errorf("error in creating file main.go: %w", err)
	}
	// f.Close() TODO
	return f, nil
}

func (p *Project) CreateStructure(nameProject string) error {

	// создаем главную папку проекта
	err := os.Mkdir(nameProject, perm)
	if err != nil {
		return fmt.Errorf("error in creating general directory: %w", err)
	}

	// добавляем к пути имя проекта (теперь в пути рабочая директория)
	func(path string) {
		count := utf8.RuneCountInString(path)
		if path[count-1] == '/' {
			p.Path += nameProject
		} else {
			p.Path += "/" + nameProject
		}
	}(p.Path)

	// изменяем рабочую директорию
	err = os.Chdir(p.Path)
	if err != nil {
		return fmt.Errorf("error in change work directory for genos: %w", err)
	}

	// инициализация модуля
	err = createGoModule(nameProject)
	if err != nil {
		return err
	}

	// mkdir cmd/main & touch main.go
	mainFile, err := createMain()
	if err != nil {
		return err
	}

	// mkdir configs
	err = os.Mkdir("configs", perm)
	if err != nil {
		return fmt.Errorf("error in creating configs directory: %w", err)
	}

	// mkdir docs
	err = os.Mkdir("docs", perm)
	if err != nil {
		return fmt.Errorf("error in creating docs directory: %w", err)
	}

	// mkdir internal
	err = os.Mkdir("internal", perm)
	if err != nil {
		return fmt.Errorf("error in creating internal directory: %w", err)
	}

	// mkdir pkg
	err = os.Mkdir("pkg", perm)
	if err != nil {
		return fmt.Errorf("error in creating pkg directory: %w", err)
	}

	commands.CreateMainCode(mainFile)

	return nil
}
