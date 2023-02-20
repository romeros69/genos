package app

import (
	"fmt"
	"genos/internal/commands"
	"log"
	"os"
)

func Run() {
	generateCommand := new(commands.Generate)
	helpCommand := new(commands.Help)
	initLayoutCommand := new(commands.InitLayout)

	countArgs := len(os.Args)
	if countArgs == 1 {
		fmt.Println("No input option\n" +
			"Use option `-h` or `--help` for see all commands")
		os.Exit(0)
	} else if countArgs > 3 {
		fmt.Println("Too many option!\n" +
			"Use option `-h` or `--help` for help")
		os.Exit(0)
	} else {
		switch {
		case (os.Args[1] == generateCommand.GetNames()[0]) || (os.Args[1] == generateCommand.GetNames()[1]):
			err := generateCommand.Do()
			if err != nil {
				log.Fatalf("error in do function generateCommand: %s", err) // FIXME - fix handle error
			}
		case (os.Args[1] == helpCommand.GetNames()[0]) || (os.Args[1] == helpCommand.GetNames()[1]):
			err := helpCommand.Do()
			if err != nil {
				log.Fatalf("error in do function helpCommand: %s", err) // FIXME - fix handle error
			}
		case (os.Args[1] == initLayoutCommand.GetNames()[0]) || (os.Args[1] == initLayoutCommand.GetNames()[1]):
			err := initLayoutCommand.Do()
			if err != nil {
				log.Fatal("error in do function initLayoutCommand: %s", err) // FIXME - fix handle error
			}
		default:
			fmt.Printf("Unknown command\n" +
				"Use option `-h` or `--help` for see all comands")
			os.Exit(0)
		}
	}

	//fmt.Printf("Please, enter full path to work directory, or press 'enter' for use this directory\n>")
	//
	//project := new(layout.Project)
	//
	//fmt.Scanln(&project.Path)
	//
	//if project.Path == "" {
	//	path, err := os.Getwd()
	//	if err != nil {
	//		log.Fatal("error in getting current path")
	//	}
	//	project.Path = path
	//} else {
	//	_, err := os.Stat(project.Path)
	//	if err != nil {
	//		log.Fatal("error in check work dir path")
	//	}
	//	if os.IsNotExist(err) {
	//		log.Fatal("error in check work dir path 2")
	//	}
	//
	//	err = os.Chdir(project.Path)
	//	if err != nil {
	//		log.Fatal("error in change work dir for genos")
	//	}
	//}
	//
	//fmt.Printf("Please, enter name of application\n>")
	//
	//var nameProject string
	//
	//fmt.Scan(&nameProject)
	//
	//err := project.CreateStructure(nameProject)
	//if err != nil {
	//	panic(err.Error())
	//}
}
