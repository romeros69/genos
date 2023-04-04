package files

import (
	"fmt"
	"genos/internal/usecase"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
)

type FileSource struct {
}

func NewFileWorker() *FileSource {
	return &FileSource{}
}

var _ usecase.FileSourceWorker = (*FileSource)(nil)

func (fw *FileSource) CreateFile(path string) (*os.File, error) {
	var err error
	file, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("error in creating file: %s: %w", path, err)
	}
	return file, nil
}

func (fw *FileSource) WriteAST(file *os.File, ast *ast.File) error {
	fset := token.NewFileSet()
	err := printer.Fprint(file, fset, ast)
	if err != nil {
		return fmt.Errorf("error in generate %s: %w", file.Name(), err)
	}
	return nil
}

func (fw *FileSource) CloseFile(file *os.File) error {
	err := file.Close()
	if err != nil {
		return fmt.Errorf("error in closing file %s: %w", file.Name(), err)
	}
	return nil
}
