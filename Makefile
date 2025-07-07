# Makefile for the Go Commit Linter

# Define the output directory for the binary.
BIN_DIR := bin

# Define the name of the binary.
BINARY_NAME := cl

# Define the Go source files.
GO_FILES := $(wildcard *.go)

# The default target, which is an alias for 'build'.
.PHONY: all
all: build

# Build the Go application and place it in the BIN_DIR.
.PHONY: build
build: $(BIN_DIR)/$(BINARY_NAME)

$(BIN_DIR)/$(BINARY_NAME): $(GO_FILES) go.mod
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $@ .
	@echo "✅ $(BINARY_NAME) built successfully in $(BIN_DIR)/"

# Clean the build artifacts by removing the BIN_DIR.
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@echo "Cleanup complete."
