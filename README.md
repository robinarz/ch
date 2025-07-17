# Commit Helper (ch)

A fast, lightweight Go tool to help you create beautiful, conventional commits.

This tool provides:

 - An interactive, step-by-step commit builder.
 - A validator to enforce Conventional Commits on messages created manually.

## Installation

You can download the latest pre-compiled binary for your operating system.

```bash
# macOS (Apple Silicon / arm64)

curl -L -o ch https://github.com/robinarz/ch/releases/latest/download/ch-darwin-arm64
chmod +x ch
sudo mv ch $HOME/.local/bin/
```

```bash
# macOS (Intel / amd64)

curl -L -o ch https://github.com/robinarz/ch/releases/latest/download/ch-darwin-amd64
chmod +x ch
sudo mv ch $HOME/.local/bin/
```

```bash
# Linux (amd64)

curl -L -o ch https://github.com/robinarz/ch/releases/latest/download/ch-linux-amd64
chmod +x ch
sudo mv ch $HOME/.local/bin/
```

## Usage
### Interactive Commit Builder (Recommended)

The primary way to use this tool is with the commit command. It will stage all your changes and guide you through creating a perfect commit message.

Simply run this instead of git commit:

```bash
ch commit
```

### Manual Commit Validation (Optional Git Hook)

If you still want to use `git commit -m "..."` manually, you can use **ch** as a safeguard to validate your message.

- Create a file at .git/hooks/commit-msg in your repository.

- Add the following script:

```bash
#!/bin/bash

# Path to your installed ch executable
LINTER_PATH="$HOME/.local/bin/ch"

# Validate the commit message file
if "$LINTER_PATH" validate "$1"; then
  exit 0
else
  exit 1
fi
```

- Make the hook executable:

```bash
chmod +x .git/hooks/commit-msg
```

### Building from Source

If you prefer to build the tool from source, you'll need Go (version 1.24 or later) and make.


- Clone the repository:

```bash
git clone https://github.com/robinarz/ch.git
cd ch
```

- Build the binary:

```bash
make build
```
This creates the executable at ./bin/ch.

- Install the binary (optional):
You can install the compiled binary to your local bin directory to make it available everywhere.

```bash
make install
```
    
