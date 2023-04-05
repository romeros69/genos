package model

import "go/ast"

type BaseGenerator interface {
	FullPathToFile() string
	GenAST() *ast.File
}
