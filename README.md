# Commit Helper

A fast, lightweight Go tool to enforce Conventional Commits as a Git hook.

## Quick Start

Prerequisites: Go (version 1.24 or later) and make.

### 1. Build the Linter

From this project's root directory, run:

```bash
make build
```
This creates the executable at bin/ch.

### 2. Install as a Git Hook

In the project repository where you want to use the linter:

Copy the Binary: Copy the bin directory (containing the commit-helper executable) into your project's root folder.

Create the Hook Script: Create a file at `.git/hooks/commit-msg` and add the following two lines:

```bash
#!/bin/bash

LINTER_PATH="$HOME/.local/bin/ch"
COMMIT_MSG_FILE="$1"

# The hook's only job is to validate the commit message file.
# The `validate` subcommand will print its own success or error messages
# to the terminal's standard error stream, which preserves colors.
# We only need to check the exit code.
if "$LINTER_PATH" validate "$COMMIT_MSG_FILE"; then
  # If the linter exits with 0, the commit is valid.
  exit 0
else
  # If the linter exits with a non-zero code, the commit is invalid.
  exit 1
fi
```

### 3. Make the Hook Executable

Make the hook executable:

```bash
chmod +x .git/hooks/commit-msg
```

That's it! Your commit messages will now be validated automatically.
