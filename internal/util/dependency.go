package util

import (
	"go/ast"
	"go/token"
	"os/exec"
)

func DownloadDependency(file *ast.File) error {
	impStr := getStringImportSlice(file)
	for _, v := range impStr {
		cmd := exec.Command("go", "get", "-u", cutQuotes(v))
		_ = cmd.Run()
	}
	return nil
}

func getStringImportSlice(file *ast.File) []string {
	specs := getSpecImportSlice(file)
	importsSlice := make([]string, len(specs))
	for i, v := range specs {
		impSpec, ok := v.(*ast.ImportSpec)
		if ok {
			importsSlice[i] = impSpec.Path.Value
		}
	}
	return importsSlice
}

func getSpecImportSlice(file *ast.File) []ast.Spec {
	for _, v := range file.Decls {
		decl, ok := v.(*ast.GenDecl)
		if ok && decl.Tok == token.IMPORT {
			return decl.Specs
		}
	}
	return nil
}

func cutQuotes(s string) string {
	runeSlice := []rune(s)
	l := len(runeSlice)
	newRune := make([]rune, 0)
	for i, v := range runeSlice {
		if i != 0 && i != l-1 {
			newRune = append(newRune, v)
		}
	}
	return string(newRune)
}
