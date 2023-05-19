%{
package dsl

const (
  INT = iota
  INT8
  INT16
  INT32
  INT64
  UINT
  UINT8
  UINT16
  UINT32
  UINT64
  BYTE
  RUNE
  FLOAT32
  FLOAT64
  BOOL
  STRING
)

const (
  CREATE = "create"
  READ 	 = "read"
  UPDATE = "update"
  DELETE = "delete"
  LIST   = "list"
)

type Field struct {
  TokType int
  Name string
}

type Ent struct {
  Name string
  Fields []Field
  Actions []string
}

type AST struct {
  Entities []Ent
}

%}

%union{
  ast AST
  entities []Ent
  ent Ent
  fields []Field
  field Field
  val string
  tokType int
  actions []string
}

%token <val>	 	IDENT EN_TOK LS_TOK RS_TOK ARROW_TOK COMMA_TOK CREATE_TOK READ_TOK UPDATE_TOK DELETE_TOK LIST_TOK
%token <val>		INT_T
%token <val>		INT8_T
%token <val>		INT16_T
%token <val>		INT32_T
%token <val>		INT64_T
%token <val>		UINT_T
%token <val>		UINT8_T
%token <val>		UINT16_T
%token <val>		UINT32_T
%token <val>		UINT64_T
%token <val>		BYTE_T
%token <val>		RUNE_T
%token <val>		FLOAT32_T
%token <val>		FLOAT64_T
%token <val>		BOOL_T
%token <val>		STRING_T

%type  <ast>     	ast
%type  <entities> 	entities
%type  <fields>	 	fields
%type  <actions> 	actions
%type  <ent>     	ent
%type  <field>   	field
%type  <val>  	 	action
%type  <tokType> 	type

%start main

%%

main: ast
  {
    yylex.(*Lex).result = $1
  }

ast: entities
  {
    $$ = AST{Entities: $1}
  }

entities:
  ent
  {
    $$ = []Ent{$1}
  }
| entities ent
  {
    $$ = append($1, $2)
  }

ent: EN_TOK IDENT LS_TOK fields RS_TOK ARROW_TOK actions
  {
    $$ = Ent{
    	Name: $2,
    	Fields: $4,
    	Actions: $7,
    }
  }

fields:
  field
  {
    $$ = []Field{$1}
  }
| fields field
  {
    $$ = append($1, $2)
  }

field: IDENT type
  {
    $$ = Field{Name: $1, TokType: $2}
  }

actions:
  action
  {
    $$ = []string{$1}
  }
| actions COMMA_TOK action
  {
    $$ = append($1, $3)
  }

action:
  CREATE_TOK
  {
    $$ = CREATE
  }
| READ_TOK
  {
    $$ = READ
  }
| UPDATE_TOK
  {
    $$ = UPDATE
  }
| DELETE_TOK
  {
    $$ = DELETE
  }
| LIST_TOK
  {
    $$ = LIST
  }


type:
  INT_T
  {
    $$ = INT
  }
| INT8_T
  {
    $$ = INT8
  }
| INT16_T
  {
    $$ = INT16
  }
| INT32_T
  {
    $$ = INT32
  }
| INT64_T
  {
    $$ = INT64
  }
| UINT_T
  {
    $$ = UINT
  }
| UINT8_T
  {
    $$ = UINT8
  }
| UINT16_T
  {
    $$ = UINT16
  }
| UINT32_T
  {
    $$ = UINT32
  }
| UINT64_T
  {
    $$ = UINT64
  }
| BYTE_T
  {
    $$ = BYTE
  }
| RUNE_T
  {
    $$ = RUNE
  }
| FLOAT32_T
  {
    $$ = FLOAT32
  }
| FLOAT64_T
  {
    $$ = FLOAT64
  }
| BOOL_T
  {
    $$ = BOOL
  }
| STRING_T
  {
    $$ = STRING
  }
