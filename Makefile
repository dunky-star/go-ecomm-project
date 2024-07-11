build:
	@go build -o bin/go-ecomm-project cmd/main

test:
	@go test -v ./...

run: build
	@.bin/go-ecomm-project
