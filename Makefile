run: clean build

build:
	go build -o bin/genos cmd/main/main.go

clean:
	rm -rf /Users/romeros/Documents/diplomchik/worktest/*