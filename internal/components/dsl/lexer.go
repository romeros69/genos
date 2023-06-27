// Code generated by golex. DO NOT EDIT.

package dsl

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

type Lex struct {
	input  []byte
	pos    int
	result AST
	err    error
}

var buf = bytes.NewBuffer(nil)

func NewLex(input []byte) *Lex {
	return &Lex{input: input}
}

func Parse(yyLex yyLexer) {
	_ = yyParse(yyLex)
}

func (l *Lex) GetResult() AST {
	return l.result
}

func (l *Lex) Next() byte {
	if l.pos >= len(l.input) || l.pos == -1 {
		if l.pos == len(l.input) {
			buf.WriteByte(l.input[l.pos-1])
		}
		l.pos = -1
		return 0
	}
	if l.pos != 0 {
		buf.WriteByte(l.input[l.pos-1])
	}
	l.pos++
	return l.input[l.pos-1]
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

yystate0:

	buf.Reset()

	goto yystart1

yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= '+' || c >= '.' && c <= '@' || c >= '[' && c <= '`' || c == '|' || c >= '~' && c <= 'ÿ'
	case c == ',':
		goto yystate6
	case c == '-':
		goto yystate7
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\r' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == 'b':
		goto yystate11
	case c == 'c':
		goto yystate18
	case c == 'd':
		goto yystate24
	case c == 'e':
		goto yystate30
	case c == 'f':
		goto yystate36
	case c == 'i':
		goto yystate45
	case c == 'l':
		goto yystate55
	case c == 'r':
		goto yystate59
	case c == 's':
		goto yystate66
	case c == 'u':
		goto yystate72
	case c == '{':
		goto yystate88
	case c == '}':
		goto yystate89
	case c >= 'A' && c <= 'Z' || c == 'a' || c == 'g' || c == 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'q' || c == 't' || c >= 'v' && c <= 'z':
		goto yystate9
	}

yystate2:
	c = l.Next()
	goto yyrule29

yystate3:
	c = l.Next()
	goto yyrule30

yystate4:
	c = l.Next()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate5:
	c = l.Next()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate5
	}

yystate6:
	c = l.Next()
	goto yyrule4

yystate7:
	c = l.Next()
	switch {
	default:
		goto yyrule30
	case c == '>':
		goto yystate8
	}

yystate8:
	c = l.Next()
	goto yyrule5

yystate9:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate10:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate11:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate12
	case c == 'y':
		goto yystate15
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'x' || c == 'z':
		goto yystate10
	}

yystate12:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate13
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate13:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate14
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate10
	}

yystate14:
	c = l.Next()
	switch {
	default:
		goto yyrule25
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate15:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate16
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate16:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate17
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate17:
	c = l.Next()
	switch {
	default:
		goto yyrule21
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate18:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'r':
		goto yystate19
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate10
	}

yystate19:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate20
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate20:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate21
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate21:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate22
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate22:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate23
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate23:
	c = l.Next()
	switch {
	default:
		goto yyrule6
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate24:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate25
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate25:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate26
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate10
	}

yystate26:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate27
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate27:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate28
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate28:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate29
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate29:
	c = l.Next()
	switch {
	default:
		goto yyrule9
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate30:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate31
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate31:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate32
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate32:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'i':
		goto yystate33
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate33:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate34
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate34:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'y':
		goto yystate35
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate10
	}

yystate35:
	c = l.Next()
	switch {
	default:
		goto yyrule27
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate36:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'l':
		goto yystate37
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate10
	}

yystate37:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'o':
		goto yystate38
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate38:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate39
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate39:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate40
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate40:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == '3':
		goto yystate41
	case c == '6':
		goto yystate43
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate41:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '2':
		goto yystate42
	}

yystate42:
	c = l.Next()
	goto yyrule23

yystate43:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '4':
		goto yystate44
	}

yystate44:
	c = l.Next()
	goto yyrule24

yystate45:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate46
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate46:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate47
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate47:
	c = l.Next()
	switch {
	default:
		goto yyrule11
	case c == '1':
		goto yystate48
	case c == '3':
		goto yystate50
	case c == '6':
		goto yystate52
	case c == '8':
		goto yystate54
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate48:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '6':
		goto yystate49
	}

yystate49:
	c = l.Next()
	goto yyrule13

yystate50:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '2':
		goto yystate51
	}

yystate51:
	c = l.Next()
	goto yyrule14

yystate52:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '4':
		goto yystate53
	}

yystate53:
	c = l.Next()
	goto yyrule15

yystate54:
	c = l.Next()
	goto yyrule12

yystate55:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'i':
		goto yystate56
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate56:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 's':
		goto yystate57
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate10
	}

yystate57:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate58
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate58:
	c = l.Next()
	switch {
	default:
		goto yyrule10
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate59:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate60
	case c == 'u':
		goto yystate63
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate10
	}

yystate60:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate61
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate61:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'd':
		goto yystate62
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate10
	}

yystate62:
	c = l.Next()
	switch {
	default:
		goto yyrule7
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate63:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate64
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate64:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate65
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate65:
	c = l.Next()
	switch {
	default:
		goto yyrule22
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate66:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate67
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate67:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'r':
		goto yystate68
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate10
	}

yystate68:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'i':
		goto yystate69
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate69:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate70
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate70:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'g':
		goto yystate71
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate10
	}

yystate71:
	c = l.Next()
	switch {
	default:
		goto yyrule26
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate72:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'i':
		goto yystate73
	case c == 'p':
		goto yystate83
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate10
	}

yystate73:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'n':
		goto yystate74
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate74:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate75
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate75:
	c = l.Next()
	switch {
	default:
		goto yyrule16
	case c == '1':
		goto yystate76
	case c == '3':
		goto yystate78
	case c == '6':
		goto yystate80
	case c == '8':
		goto yystate82
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate76:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '6':
		goto yystate77
	}

yystate77:
	c = l.Next()
	goto yyrule18

yystate78:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '2':
		goto yystate79
	}

yystate79:
	c = l.Next()
	goto yyrule19

yystate80:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '4':
		goto yystate81
	}

yystate81:
	c = l.Next()
	goto yyrule20

yystate82:
	c = l.Next()
	goto yyrule17

yystate83:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'd':
		goto yystate84
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate10
	}

yystate84:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'a':
		goto yystate85
	case c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate85:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 't':
		goto yystate86
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate86:
	c = l.Next()
	switch {
	default:
		goto yyrule28
	case c == 'e':
		goto yystate87
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate87:
	c = l.Next()
	switch {
	default:
		goto yyrule8
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate88:
	c = l.Next()
	goto yyrule2

yystate89:
	c = l.Next()
	goto yyrule3

yyrule1: // {C}
	{
		// continue
		goto yystate0
	}
yyrule2: // {LS}
	{
		{
			l.Backup()
			return LS_TOK
		}
		goto yystate0
	}
yyrule3: // {RS}
	{
		{
			l.Backup()
			return RS_TOK
		}
		goto yystate0
	}
yyrule4: // {COMMA}
	{
		{
			l.Backup()
			return COMMA_TOK
		}
		goto yystate0
	}
yyrule5: // {ARROW}
	{
		{
			l.Backup()
			return ARROW_TOK
		}
		goto yystate0
	}
yyrule6: // {CREATE}
	{
		{
			l.Backup()
			return CREATE_TOK
		}
		goto yystate0
	}
yyrule7: // {READ}
	{
		{
			l.Backup()
			return READ_TOK
		}
		goto yystate0
	}
yyrule8: // {UPDATE}
	{
		{
			l.Backup()
			return UPDATE_TOK
		}
		goto yystate0
	}
yyrule9: // {DELETE}
	{
		{
			l.Backup()
			return DELETE_TOK
		}
		goto yystate0
	}
yyrule10: // {LIST}
	{
		{
			l.Backup()
			return LIST_TOK
		}
		goto yystate0
	}
yyrule11: // {INT_T}
	{
		{
			l.Backup()
			return INT_T
		}
		goto yystate0
	}
yyrule12: // {INT8_T}
	{
		{
			l.Backup()
			return INT8_T
		}
		goto yystate0
	}
yyrule13: // {INT16_T}
	{
		{
			l.Backup()
			return INT16_T
		}
		goto yystate0
	}
yyrule14: // {INT32_T}
	{
		{
			l.Backup()
			return INT32_T
		}
		goto yystate0
	}
yyrule15: // {INT64_T}
	{
		{
			l.Backup()
			return INT64_T
		}
		goto yystate0
	}
yyrule16: // {UINT_T}
	{
		{
			l.Backup()
			return UINT_T
		}
		goto yystate0
	}
yyrule17: // {UINT8_T}
	{
		{
			l.Backup()
			return UINT8_T
		}
		goto yystate0
	}
yyrule18: // {UINT16_T}
	{
		{
			l.Backup()
			return UINT16_T
		}
		goto yystate0
	}
yyrule19: // {UINT32_T}
	{
		{
			l.Backup()
			return UINT32_T
		}
		goto yystate0
	}
yyrule20: // {UINT64_T}
	{
		{
			l.Backup()
			return UINT64_T
		}
		goto yystate0
	}
yyrule21: // {BYTE_T}
	{
		{
			l.Backup()
			return BYTE_T
		}
		goto yystate0
	}
yyrule22: // {RUNE_T}
	{
		{
			l.Backup()
			return RUNE_T
		}
		goto yystate0
	}
yyrule23: // {FLOAT32_T}
	{
		{
			l.Backup()
			return FLOAT32_T
		}
		goto yystate0
	}
yyrule24: // {FLOAT64_T}
	{
		{
			l.Backup()
			return FLOAT64_T
		}
		goto yystate0
	}
yyrule25: // {BOOL_T}
	{
		{
			l.Backup()
			return BOOL_T
		}
		goto yystate0
	}
yyrule26: // {STRING_T}
	{
		{
			l.Backup()
			return STRING_T
		}
		goto yystate0
	}
yyrule27: // {ENTITY}
	{
		{
			l.Backup()
			return EN_TOK
		}
		goto yystate0
	}
yyrule28: // {IDENT}
	{
		{
			l.Backup()
			lval.val = buf.String()
			return IDENT
		}
		goto yystate0
	}
yyrule29: // \0
	{
		{
			return 0
		} // Exit on EOF or any other error
		goto yystate0
	}
yyrule30: // .
	if true { // avoid go vet determining the below panic will not be reached
		{
			fmt.Printf("ERROR 2: %s, buf: %s\n", string(c), buf)
			return -1
		}
		goto yystate0
	}
	panic("unreachable")

yyabort: // no lexem recognized
	// silence unused label errors for build and satisfy go vet reachability analysis
	{
		if false {
			goto yyabort
		}
		if false {
			goto yystate0
		}
		if false {
			goto yystate1
		}
	}

	fmt.Println("ERROR 3")
	log.Fatal("scanner internal error")
	return 0
}