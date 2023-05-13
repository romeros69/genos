package service

import (
	"go/ast"
	"io/fs"
	"os"
)

type FileSourceWorker interface {
	WriteAST(*os.File, *ast.File) error
	CreateFile(string) (*os.File, error)
	OpenFile(string) (*os.File, error)
	ReadFile(*os.File) ([]byte, error)
	CloseFile(*os.File) error
	ReadAST(string) (*ast.File, error)
	WriteByteSliceFile(*os.File, []byte) error
}

type FolderSourceWorker interface {
	CreateDir(string, fs.FileMode) error
}

type ExecuteCLI interface {
	ExecuteCommand(string, ...string) error
}

type HelpContract interface {
	Help() string
}

type InitLayoutContract interface {
	InitLayoutDo(string, string) error
}

type CliCommandContract interface {
	Format(string) error
	CreateGoModule(string) error
}

type FolderContract interface {
	CreateFolder(string, string) error
}

type GenerateContract interface {
	GenerateDo(string, string) error
}
