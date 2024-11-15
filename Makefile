# Name of the output binary
BINARY_NAME=temporal-phantom-worker
BIN_DIR=bin
RELEASE_DIR=release

# Go settings
GOENVVARS=CGO_ENABLED=0
GO=go
GOFLAGS=

# Create directories in a cross-platform manner
MKDIR = mkdir -p

# Default release version if not specified
VERSION ?= $(shell git describe --tags --abbrev=0)

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: build
build: build-linux build-windows build-darwin ## Build the project binaries for all supported platforms

# Cross-compilation targets
.PHONY: build-linux
build-linux: ## Build the project binary for Linux
	@ echo "Building for Linux"
	@ $(MKDIR) $(BIN_DIR)/linux
	@ GOOS=linux GOARCH=amd64 $(GOENVVARS) $(GO) build $(GOFLAGS) -o $(BIN_DIR)/linux/$(BINARY_NAME) .

.PHONY: build-windows
build-windows: ## Build the project binary for Windows
	@ echo "Building for Windows"
	@ $(MKDIR) $(BIN_DIR)/windows
	@ GOOS=windows GOARCH=amd64 $(GOENVVARS) $(GO) build $(GOFLAGS) -o $(BIN_DIR)/windows/$(BINARY_NAME).exe .

.PHONY: build-darwin
build-darwin: ## Build the project binary for Darwin
	@ echo "Building for Darwin"
	@ $(MKDIR) $(BIN_DIR)/darwin
	@ GOOS=darwin GOARCH=amd64 $(GOENVVARS) $(GO) build $(GOFLAGS) -o $(BIN_DIR)/darwin/$(BINARY_NAME) .

.PHONY: test
test: ## Run all tess with verbose output
	@ echo "Running tests"
	@ $(GO) test ./... -v

# Clean up generated files
.PHONY: clean
clean: ## Clean all generated files
	@ echo "Cleaning bin"
	@ rm -rf $(BIN_DIR)
	@ echo "Cleaning release"
	@ rm -rf $(RELEASE_DIR)

.PHONY: release
release: release-linux release-windows release-darwin ## Create release packages for all platforms

# Package binaries into tar.gz files with versioning
.PHONY: release-linux
release-linux: build-linux ## Create release package for Linux
	@ echo "Creating Linux release package"
	@ mkdir -p $(RELEASE_DIR)
	@ tar -C $(BIN_DIR)/linux -cvf $(RELEASE_DIR)/$(BINARY_NAME)-linux-amd64-$(VERSION).tar.gz $(BINARY_NAME)

.PHONY: release-windows
release-windows: build-windows ## Create release package for Windows
	@ echo "Creating Windows release package"
	@ mkdir -p $(RELEASE_DIR)
	@ tar -C $(BIN_DIR)/windows -cvf $(RELEASE_DIR)/$(BINARY_NAME)-windows-amd64-$(VERSION).tar.gz $(BINARY_NAME).exe

.PHONY: release-darwin
release-darwin: build-darwin ## Create release package for Darwin
	@ echo "Creating Darwin release package"
	@ mkdir -p $(RELEASE_DIR)
	@ tar -C $(BIN_DIR)/darwin -cvf $(RELEASE_DIR)/$(BINARY_NAME)-darwin-amd64-$(VERSION).tar.gz $(BINARY_NAME)

# Run both build and tests
.PHONY: all
all: clean test build release