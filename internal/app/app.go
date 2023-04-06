package app

import (
	"fmt"
	"genos/internal/command"
	"genos/internal/infrastructure/cli"
	"genos/internal/infrastructure/files"
	"genos/internal/infrastructure/folders"
	"genos/internal/service"
	"log"
	"os"
)

func Run() {
	helpUC := service.NewHelpUseCase()
	cliUC := service.NewCliUC(cli.NewExecuteSLI())
	folderUC := service.NewFolderUC(folders.NewFolderSource())
	initLayoutUC := service.NewInitLayout(files.NewFileSource(), cliUC, folderUC)

	initLayoutCommand := command.NewInitLayout(initLayoutUC)
	helpCommand := command.NewHelp(helpUC)

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
		case (os.Args[1] == helpCommand.GetNames()[0]) || (os.Args[1] == helpCommand.GetNames()[1]):
			err := helpCommand.Do()
			if err != nil {
				log.Fatalf("error in do function helpCommand: %s", err) // FIXME - fix handle error
			}
		case (os.Args[1] == initLayoutCommand.GetNames()[0]) || (os.Args[1] == initLayoutCommand.GetNames()[1]):
			err := initLayoutCommand.Do()
			if err != nil {
				log.Fatalf("error in do function initLayoutCommand: %s", err) // FIXME - fix handle error
			}
		default:
			fmt.Printf("Unknown command\n" +
				"Use option `-h` or `--help` for see all comands")
			os.Exit(0)
		}
	}
}
