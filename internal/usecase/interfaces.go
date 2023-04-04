package usecase

import (
	"go/ast"
	"os"
)

type FileStorage interface {
	WriteAST(*os.File, *ast.File) error
	CreateFile(string) (*os.File, error)
	CloseFile(*os.File) error
	// TODO ReadAST(f *os.File) (*ast.File, error)
}
