# Makefile for the Go Commit Linter

# Define the output directory for the binary.
BIN_DIR := bin

# Define the name and path of the binary.
BINARY_NAME := ch
CMD_PATH := ./cmd/$(BINARY_NAME)

# The default target, which is an alias for 'build'.
.PHONY: all
all: build

# Build the Go linter application.
.PHONY: build
build: $(BIN_DIR)/$(BINARY_NAME)

$(BIN_DIR)/$(BINARY_NAME):
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $@ $(CMD_PATH)
	@echo "✅ $(BINARY_NAME) built successfully in $(BIN_DIR)/"

# Clean the build artifacts by removing the BIN_DIR.
.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)
	@echo "Cleanup complete."

# This makes the 'ch' command available globally for the current user.
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME) to $(HOME)/.local/bin..."
	@mkdir -p $(HOME)/.local/bin
	@cp $(BIN_DIR)/$(BINARY_NAME) $(HOME)/.local/bin/
	@echo "✅ $(BINARY_NAME) installed successfully. Make sure $(HOME)/.local/bin is in your PATH."
