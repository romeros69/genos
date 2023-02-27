package commands

import (
	"fmt"
	"genos/internal/gen"
	"io/fs"
	"os"
	"os/exec"
	"unicode/utf8"
)

const perm fs.FileMode = 0777

type InitLayout struct {
	Path        string
	NameProject string
}

// Функция установки пути к рабочей директории
// Путь получается от пользователя
// Или берется текущая, откужа была запущена программа
func (i *InitLayout) readWorkDirectory() error {
	fmt.Printf("Enter full path to work directory, or press `enter` for use this directory\n")
	_, err := fmt.Scanln(&i.Path)
	if err != nil {
		return fmt.Errorf("error in scan path in readWorkDirectory(): %w", err)
	}
	if i.Path == "" {
		path, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error in get current work directory readWorkDirectory(): %w", err)
		}
		i.Path = path
	} else {
		_, err := os.Stat(i.Path)
		if err != nil {
			return fmt.Errorf("error in get stat input work directory readWorkDirectory(): %w", err)
		}
		if os.IsNotExist(err) {
			return fmt.Errorf("error in get stat input work directory readWorkDirectory(): %w", err)
		}
		err = os.Chdir(i.Path)
		if err != nil {
			return fmt.Errorf("error in change work directory readWorkDirectory(): %w", err)
		}
	}
	return nil
}

// Функция установки именя проекта
func (i *InitLayout) readNameProject() error {
	fmt.Println("Enter project name")
	_, err := fmt.Scan(&i.NameProject)
	if err != nil {
		return fmt.Errorf("error in get name of project readNameProject: %w", err)
	}
	return nil
}

// Форматирование пути - добавление к нему имени проекта
func (i *InitLayout) formatFullPath() {
	count := utf8.RuneCountInString(i.Path)
	if i.Path[count-1] == '/' {
		i.Path += i.NameProject
	} else {
		i.Path += "/" + i.NameProject
	}
}

// функция инициализации go-модуля
func (i *InitLayout) createGoModule() error {
	cmd := exec.Command("go", "mod", "init", i.NameProject)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error in execute go mod init: %w", err)
	}
	return nil
}

// функция генерации структуры файловых директорий
func (i *InitLayout) createFileStructure() error {
	// создание директории проекта
	err := os.Mkdir(i.NameProject, perm)
	if err != nil {
		return fmt.Errorf("error in creating general directory: %w", err)
	}

	// изменяем текущую рабучую директорию касательно генератора
	err = os.Chdir(i.Path)
	if err != nil {
		return fmt.Errorf("error in change work dir createFileStructure: %w", err)
	}

	// mkdir cmd/main
	err = os.MkdirAll("cmd/main", perm)
	if err != nil {
		return fmt.Errorf("error in creating cmd & main dir: %w", err)
	}

	// mkdir config
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

	// mkdir app
	err = os.Mkdir("internal/app", perm)
	if err != nil {
		return fmt.Errorf("error in creating app directory: %w", err)
	}

	// mkdir controller
	err = os.Mkdir("internal/controller", perm)
	if err != nil {
		return fmt.Errorf("error in creating controller directory: %w", err)
	}

	// mkdir entity
	err = os.Mkdir("internal/entity", perm)
	if err != nil {
		return fmt.Errorf("error in creating entity directory: %w", err)
	}

	// mkdir usecase
	err = os.Mkdir("internal/usecase", perm)
	if err != nil {
		return fmt.Errorf("error in creating usecase directory: %w", err)
	}

	// mkdir repo
	err = os.Mkdir("internal/usecase/repo", perm)
	if err != nil {
		return fmt.Errorf("error in creating repo directory: %w", err)
	}

	// mkdir pkg
	err = os.Mkdir("pkg", perm)
	if err != nil {
		return fmt.Errorf("error in creating pkg directory: %w", err)
	}

	// mkdir httpserver
	err = os.Mkdir("pkg/httpserver", perm)
	if err != nil {
		return fmt.Errorf("error in creating httpserver directory: %w", err)
	}

	// mkdir postgres
	err = os.Mkdir("pkg/postgres", perm)
	if err != nil {
		return fmt.Errorf("error in creating postgres directory: %w", err)
	}

	return nil
}

// Do - TODO
// При вызове ее мы находимся в path
func (i *InitLayout) Do() error {
	err := i.readWorkDirectory()
	if err != nil {
		return fmt.Errorf("init layout - do: %w", err)
	}
	err = i.readNameProject()
	if err != nil {
		return fmt.Errorf("init layout - do: %w", err)
	}
	i.formatFullPath()
	err = i.createFileStructure()
	if err != nil {
		return fmt.Errorf("init layout - do: %w", err)
	}
	// инициализируем модуль
	err = i.createGoModule()
	if err != nil {
		return fmt.Errorf("init layout - do: %w", err)
	}
	err = generate.GenMain(i.NameProject)
	if err != nil {
		return fmt.Errorf("error in GenMain(): %w", err)
	}
	err = generate.GenApp()
	if err != nil {
		return fmt.Errorf("error in GenApp(): %w", err)
	}
	err = generate.GenHttpServer()
	if err != nil {
		return fmt.Errorf("error in GenHttpServer(): %w", err)
	}
	err = generate.GenOptionsHttpServer()
	if err != nil {
		return fmt.Errorf("error in GenOptionsHttpServer(): %w", err)
	}
	return nil
}

func (i *InitLayout) GetNames() []string {
	return []string{"-i", "--init"}
}
