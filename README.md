# Commit Linter

A fast, lightweight Go tool to enforce Conventional Commits as a Git hook.

## Quick Start

Prerequisites: Go (version 1.24 or later) and make.

### 1. Build the Linter

From this project's root directory, run:

```bash
make build
```
This creates the executable at bin/cl.

### 2. Install as a Git Hook

In the project repository where you want to use the linter:

Copy the Binary: Copy the bin directory (containing the commit-linter executable) into your project's root folder.

Create the Hook Script: Create a file at `.git/hooks/commit-msg` and add the following two lines:

```bash
#!/bin/bash

# Path to your compiled linter executable, relative to the repo root
LINTER_PATH="./bin/cl"

# Execute the linter
"$LINTER_PATH" "$1"

# Exit with the linter's exit code
exit $?
```

### 3. Make the Hook Executable

Make the hook executable:

```bash
chmod +x .git/hooks/commit-msg
```

That's it! Your commit messages will now be validated automatically.
