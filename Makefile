build:
	@go build -o bin/go-ecomm-project cmd/main

test:
	@go test -v ./...

run: build
	@.bin/go-ecomm-project

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down