package files

import (
	"fmt"
	"genos/internal/service"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"os"
)

type FileSource struct {
}

func NewFileSource() *FileSource {
	return &FileSource{}
}

var _ service.FileSourceWorker = (*FileSource)(nil)

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

func (fw *FileSource) OpenFile(fullPathToDSLFile string) (*os.File, error) {
	file, err := os.Open(fullPathToDSLFile)
	if err != nil {
		return nil, fmt.Errorf("error in open file %s: %w", fullPathToDSLFile, err)
	}
	return file, nil
}

func (fw *FileSource) ReadFile(file *os.File) ([]byte, error) {
	inputBuf := make([]byte, func(f *os.File) int64 {
		stat, _ := f.Stat()
		return stat.Size()
	}(file))

	_, err := file.Read(inputBuf)
	if err == io.EOF {
		// maybe...
	} else if err != nil {
		return nil, fmt.Errorf("error in read file %s: %w", file.Name(), err)
	}
	return inputBuf, nil
}
