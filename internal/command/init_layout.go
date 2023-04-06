package command

import (
	"fmt"
	"genos/internal/service"
	"os"
	"unicode/utf8"
)

type InitLayout struct {
	ilUC        service.InitLayoutContract
	path        string
	nameProject string
}

func NewInitLayout(ilUC service.InitLayoutContract) *InitLayout {
	return &InitLayout{ilUC: ilUC}
}

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
	err = i.ilUC.InitLayoutDo(i.nameProject, i.path)
	if err != nil {
		return fmt.Errorf("init layout - do: %w", err)
	}
	return nil
}

// Функция установки пути к рабочей директории
// Путь получается от пользователя
// Или берется текущая, откужа была запущена программа
func (i *InitLayout) readWorkDirectory() error {
	fmt.Printf("Enter full path to work directory, or press `enter` for use this directory\n")
	_, err := fmt.Scanln(&i.path)
	if err != nil {
		return fmt.Errorf("error in scan path in readWorkDirectory(): %w", err)
	}
	if i.path == "" {
		path, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error in get current work directory readWorkDirectory(): %w", err)
		}
		i.path = path
	} else {
		_, err := os.Stat(i.path)
		if err != nil {
			return fmt.Errorf("error in get stat input work directory readWorkDirectory(): %w", err)
		}
		if os.IsNotExist(err) {
			return fmt.Errorf("error in get stat input work directory readWorkDirectory(): %w", err)
		}
		err = os.Chdir(i.path)
		if err != nil {
			return fmt.Errorf("error in change work directory readWorkDirectory(): %w", err)
		}
	}
	return nil
}

// Функция установки именя проекта
func (i *InitLayout) readNameProject() error {
	fmt.Println("Enter project name")
	_, err := fmt.Scan(&i.nameProject)
	if err != nil {
		return fmt.Errorf("error in get name of project readNameProject: %w", err)
	}

	return nil
}

// Форматирование пути - добавление к нему имени проекта
func (i *InitLayout) formatFullPath() {
	count := utf8.RuneCountInString(i.path)
	if i.path[count-1] == '/' {
		i.path += i.nameProject
	} else {
		i.path += "/" + i.nameProject
	}
}

func (i *InitLayout) GetNames() []string {
	return []string{"-i", "--init"}
}
