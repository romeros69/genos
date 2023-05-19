package util

import (
	"genos/internal/components/dsl"
	"go/token"
)

// TypesMap - Отображение токенов лексера на их строковые представления
var TypesMap = map[int]string{
	dsl.INT:     "int",
	dsl.INT8:    "int8",
	dsl.INT16:   "int16",
	dsl.INT32:   "int32",
	dsl.INT64:   "int64",
	dsl.UINT:    "uint",
	dsl.UINT8:   "uint8",
	dsl.UINT16:  "uint16",
	dsl.UINT32:  "uint32",
	dsl.UINT64:  "uint64",
	dsl.BYTE:    "byte",
	dsl.RUNE:    "rune",
	dsl.FLOAT32: "float32",
	dsl.FLOAT64: "float64",
	dsl.BOOL:    "bool",
	dsl.STRING:  "string",
}

func GetDefaultValueOfType(typeName string) string {
	switch typeName {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "unit64", "rune", "byte", "float32", "float64":
		return "0"
	case "bool":
		return "false"
	case "string":
		return "\"\""
	}
	return ""
}

func GetGoTokenByLexerToken(lexerTok int) token.Token {
	switch lexerTok {
	case dsl.INT, dsl.INT8, dsl.INT16, dsl.INT32, dsl.INT64, dsl.UINT, dsl.UINT8, dsl.UINT16, dsl.UINT32, dsl.UINT64, dsl.RUNE, dsl.BYTE:
		return token.INT
	case dsl.FLOAT32, dsl.FLOAT64:
		return token.FLOAT
	case dsl.BOOL:
		return -1
	case dsl.STRING:
		return token.STRING
	}
	return 0
}
