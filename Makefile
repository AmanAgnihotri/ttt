.DEFAULT_GOAL := build

BUILD_PATH = ./cmd/main
BINARY_PATH := ./bin/ttt

.PHONY: brew tidy lint test build run clean compose watch help

HOMEBREW = https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh
PACKAGES = go golangci-lint docker docker-compose colima pre-commit

# setup the dependencies via Homebrew
brew:
	@echo "Setting up development environment..."
	@which brew >/dev/null || /bin/bash -c "$(curl -fsSL $(HOMEBREW))"
	@brew update
	@for pkg in $(PACKAGES); do \
		if ! brew list | grep -qx $$pkg; then \
			echo "Installing $$pkg"; \
			brew install $$pkg --force; \
		else \
			echo "$$pkg is already installed."; \
		fi; \
	done
	@echo "Setup complete!"

# tidy for managing dependencies
tidy:
	go mod tidy

# golangci-lint for comprehensive linting, with automatic fixes where applicable
lint: tidy
	golangci-lint run --fix

# run the tests
test: lint
	go test ./...

# build the binary and place it in the bin directory
build: test
	go build -o $(BINARY_PATH) $(BUILD_PATH)

# run the binary
run: build
	@cd $(dir $(BINARY_PATH)) && ./$(notdir $(BINARY_PATH))

# remove the binary
clean:
	rm -f $(BINARY_PATH)

# run via docker compose
compose: test
	@colima status | grep -q "Running" || colima start
	docker compose up --build

# see the running container's logs
logs:
	docker compose logs -f

# help with the commands
help:
	@echo "Makefile commands:"
	@echo "  make brew      - Ensure the software dependencies for development"
	@echo "  make build     - Build the binary"
	@echo "  make run       - Run the binary"
	@echo "  make clean     - Remove the binary"
	@echo "  make compose   - Run via docker compose"
	@echo "  make logs      - See the docker container's logs"
	@echo "  make lint      - Run linters with --fix flag for automatic fixes"
	@echo "  make test      - Run the tests"
