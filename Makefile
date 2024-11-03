# Name of the output binary
BINARY_NAME=temporal-phantom-worker
BINARY_DIR=bin

# Go settings
GO=go
GOFLAGS=

# Create directories in a cross-platform manner
MKDIR = mkdir -p

# Help message
help:
	@echo "Makefile commands:"
	@echo "  make build          - Build the project into a binary."
	@echo "  make test           - Run all tests with verbose output."
	@echo "  make clean          - Remove the generated binary."
	@echo "  make all            - Build the project and run tests."
	@echo "  make help           - Display this help message."
	@echo "  make build-linux    - Build for Linux."
	@echo "  make build-windows  - Build for Windows."
	@echo "  make build-macos    - Build for macOS."

# Build the project
build: build-linux build-windows build-macos

# Cross-compilation targets
build-linux:
# 	echo "Makefile commands"
	@ $(MKDIR) $(BINARY_DIR)/linux
	@ GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BINARY_DIR)/linux/$(BINARY_NAME) .

build-windows:
	@ $(MKDIR) $(BINARY_DIR)/windows
	@ GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BINARY_DIR)/windows/$(BINARY_NAME).exe .

build-macos:
	@ $(MKDIR) $(BINARY_DIR)/macos
	@ GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) -o $(BINARY_DIR)/macos/$(BINARY_NAME) .


# Run tests
test:
	@ $(GO) test ./... -v

# Clean up generated files
clean:
	@ rm -rf $(BINARY_DIR)

# Run both build and tests
all: test build

# Build for release (same as build target but organized for clarity)
build-release: build

.PHONY: all build test clean help build-linux build-windows build-macos build-release
