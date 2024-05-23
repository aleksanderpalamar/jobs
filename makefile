.PHONY: default run build test docs clean

# Variables
PROJECT_NAME = "jobs-api"
VERSION = "0.0.1"

# Tasks
default: run-with-docs

run:
	@go run main.go

run-with-docs:
	@swag init
	@go run -tags=docs main.go

build:
	@go build -o $(PROJECT_NAME) main.go

test:
	@go test ./...

docs:
	@swag init

clean:
	@rm -rf $(PROJECT_NAME)
	@rm -rf ./docs