package app

import (
	"fmt"
	"genos/internal/layout"
	"log"
	"os"
)

func Run() {

	fmt.Printf("Please, enter full path to work directory, or press 'enter' for use this directory\n>")

	project := new(layout.Project)

	fmt.Scanln(&project.Path)

	if project.Path == "" {
		path, err  := os.Getwd()
		if err != nil {
			log.Fatal("error in getting current path")
		}
		project.Path = path
	} else {
		_, err := os.Stat(project.Path)
		if err != nil {
			log.Fatal("error in check work dir path")
		}
		if os.IsNotExist(err) {
			log.Fatal("error in check work dir path 2")
		}

		err = os.Chdir(project.Path)
		if err != nil {
			log.Fatal("error in change work dir for genos")
		}
	}

	fmt.Printf("Please, enter name of application\n>")

	fmt.Scan(&project.Name)

	err := project.CreateStructure()
	if err != nil {
		panic(err.Error())
	}
}
