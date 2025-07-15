package main

import (
	"fmt"
	"os"
)

func main() {
	// If no arguments are provided, or "interactive" is passed, run the builder.
	if len(os.Args) == 1 || (len(os.Args) > 1 && os.Args[1] == "interactive") {
		if err := runInteractiveMode(""); err != nil {
			fmt.Println(errorStyle.Render("\n[ERROR]"), err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if len(os.Args) == 2 && os.Args[1] != "interactive" {
		msgFilePath := os.Args[1]
		if err := runValidationMode(msgFilePath); err != nil {
			os.Exit(1)
		}
		fmt.Println(successStyle.Render("[SUCCESS]"), "Commit message is valid.")
		os.Exit(0)
	}

	// If arguments are incorrect, show usage.
	fmt.Println(errorStyle.Render("Invalid usage."))
	fmt.Println("To create a commit interactively: go run .")
	fmt.Println("To validate a commit file:       go run . <path_to_commit_message_file>")
	os.Exit(1)
}
