.DEFAULT_GOAL := build

BUILD_PATH = ./cmd/main
BINARY_PATH := ./bin/main

.PHONY: tidy lint build clean compose

# tidy for managing dependencies
tidy:
	go mod tidy

# golangci-lint for comprehensive linting, with automatic fixes where applicable
lint: tidy
	golangci-lint run --fix

# build the project and place the binary in the bin/ directory
build: lint
	go build -o $(BINARY_PATH) $(BUILD_PATH)

# clean up the bin directory
clean:
	rm -rf bin/

# run via docker compose
compose:
	docker compose up --build -d
	docker compose watch
