package util

import (
	"genos/internal/domain/dsl"
)

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
