package commands

import (
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

// Команда для генерации кода (genos -g --generate)

func CreateMainCode(file *os.File) {
	src := `package main
import "fmt"
func main() {
	fmt.Println("hello, it is genos")
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatalf("error in creating main code: %w", err)
	}

	err = printer.Fprint(os.Stdout, fset, f)
	if err != nil {
		log.Fatalf("error in print to main.go ast: %w", err)
	}
}
