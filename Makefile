run: clean build

build:
	go build -o bin/genos cmd/main/main.go

clean:
	rm -rf /Users/romeros/Documents/diplomchik/worktest/*

compile_dsl: lex yacc

lex:
	golex -o ./internal/domain/dsl/lexer.go ./internal/domain/dsl/lexer.l

yacc:
	goyacc -o ./internal/domain/dsl/parser.go ./internal/domain/dsl/parser.y