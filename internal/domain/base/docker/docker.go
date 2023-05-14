package docker

func GetDockerFile() []byte {
	return []byte("# Step 1: Modules caching\nFROM golang:1.18-alpine as modules\nCOPY go.mod go.sum /modules/\n" +
		"WORKDIR /modules\nRUN go mod download\n\n# Step 2: Builder\nFROM golang:1.18-alpine as builder\n" +
		"COPY --from=modules /go/pkg /go/pkg\nCOPY . /app\nWORKDIR /app\nRUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \\\n" +
		"    go build -o /bin/app ./cmd/main\n\n# Step 3: Final\nFROM scratch\nCOPY --from=builder /app/configs /configs\n" +
		"COPY --from=builder /bin/app /app\nCMD [\"/app\"]")
}

func GetDockerComposeFile() []byte {
	return []byte("version: '3.7'\n\nservices:\n  psql:\n    image: postgres\n    container_name: 'postgra'\n" +
		"    ports:\n      - \"5432:5432\"\n    environment:\n      - POSTGRES_USER=postgres\n" +
		"      - POSTGRES_PASSWORD=postgres\n      - POSTGRES_DB=postgres\n\n  app:\n    container_name: 'my-app'\n" +
		"    build:\n      context: .\n      dockerfile: Dockerfile\n    ports:\n      - \"9000:9000\"\n" +
		"    depends_on:\n      - psql")
}
