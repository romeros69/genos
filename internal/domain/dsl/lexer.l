%{
package dsl

import (
    "log"
    "bytes"
    "errors"
    "fmt"
)

type Lex struct {
    input []byte
    pos int
    result AST
    err error
}

var buf = bytes.NewBuffer(nil)

func NewLex(input []byte) *Lex {
    return &Lex{input: input}
}

func (l *Lex) Next() byte {
    if l.pos >= len(l.input) || l.pos == -1 {
        if l.pos == len(l.input) {
            buf.WriteByte(l.input[l.pos - 1])
        }
        l.pos = -1
        return 0
    }
    if l.pos != 0 {
        buf.WriteByte(l.input[l.pos - 1])
    }
    l.pos++
    return l.input[l.pos - 1]
}

func (l *Lex) Error(s string) {
    fmt.Println(s)
    l.err = errors.New(s)
}

func (l *Lex) Backup() {
    if l.pos == -1 {
        return
    }
    l.pos--
}

func (l *Lex) Lex(lval *yySymType) int {
    c := l.Next() // init
%}

%yyc c
%yyn c = l.Next()

INT_T       "int"
INT8_T      "int8"
INT16_T     "int16"
INT32_T     "int32"
INT64_T     "int64"
UINT_T      "uint"
UINT8_T     "uint8"
UINT16_T    "uint16"
UINT32_T    "uint32"
UINT64_T    "uint64"
BYTE_T      "byte"
RUNE_T      "rune"
FLOAT32_T   "float32"
FLOAT64_T   "float64"
BOOL_T      "bool"
STRING_T    "string"
ENTITY      "entity"
CREATE      "create"
READ        "read"
UPDATE      "update"
DELETE      "delete"
LIST        "list"
IDENT       [a-zA-Z]+
C           [ \t\n\r]+
LS          [{]
RS          [}]
ARROW       "->"
COMMA       [\,]

%%
    buf.Reset()

{C}                 // continue

{LS}                {
                        l.Backup()
                        return LS_TOK
                    }

{RS}                {
                        l.Backup()
                        return RS_TOK
                    }

{COMMA}             {
                        l.Backup()
                        return COMMA_TOK
                    }

{ARROW}             {
                        l.Backup()
                        return ARROW_TOK
                    }

{CREATE}            {
                        l.Backup()
                        return CREATE_TOK
                    }

{READ}              {
                        l.Backup()
                        return READ_TOK
                    }

{UPDATE}            {
                        l.Backup()
                        return UPDATE_TOK
                    }

{DELETE}            {
                        l.Backup()
                        return DELETE_TOK
                    }

{LIST}              {
                        l.Backup()
                        return LIST_TOK
                    }

{INT_T}             {
                        l.Backup()
                        return INT_T
                    }

{INT8_T}            {
                        l.Backup()
                        return INT8_T
                    }

{INT16_T}           {
                        l.Backup()
                        return INT16_T
                    }

{INT32_T}           {
                        l.Backup()
                        return INT32_T
                    }

{INT64_T}           {
                        l.Backup()
                        return INT64_T
                    }

{UINT_T}            {
                        l.Backup()
                        return UINT_T
                    }

{UINT8_T}           {
                        l.Backup()
                        return UINT8_T
                    }

{UINT16_T}          {
                        l.Backup()
                        return UINT16_T
                    }

{UINT32_T}          {
                        l.Backup()
                        return UINT32_T
                    }

{UINT64_T}          {
                        l.Backup()
                        return UINT64_T
                    }

{BYTE_T}            {
                        l.Backup()
                        return BYTE_T
                    }

{RUNE_T}            {
                        l.Backup()
                        return RUNE_T
                    }

{FLOAT32_T}         {
                        l.Backup()
                        return FLOAT32_T
                    }

{FLOAT64_T}         {
                        l.Backup()
                        return FLOAT64_T
                    }

{BOOL_T}            {
                        l.Backup()
                        return BOOL_T
                    }

{STRING_T}          {
                        l.Backup()
                        return STRING_T
                    }

{ENTITY}            {
                        l.Backup()
                        return EN_TOK
                    }

{IDENT}             {
                        l.Backup()
                        lval.val = buf.String()
                        return IDENT
                    }

\0                  {
                        return 0
                    } // Exit on EOF or any other error

.                   {
                        fmt.Printf("ERROR 2: %s, buf: %s\n", string(c), buf)
                        return -1
                    }

%%
    fmt.Println("ERROR 3")
    log.Fatal("scanner internal error")
    return 0
}