package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

// Define styles used in this file
var (
	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")).
			Bold(true)
	commandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")).
			Bold(true)
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	switch subcommand {
	case "commit":
		if err := commit(); err != nil {
			// Don't print an error if the user just cancelled the interactive menu
			if err.Error() != "interactive mode cancelled by user" {
				fmt.Fprintln(os.Stderr, errorStyle.Render("\n[ERROR]"), err)
			}
			os.Exit(1)
		}
	case "validate":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, errorStyle.Render("Usage: ch validate <path_to_commit_message_file>"))
			os.Exit(1)
		}
		if err := runValidationMode(os.Args[2]); err != nil {
			os.Exit(1)
		}
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintln(os.Stderr, titleStyle.Render("Git Helper Tool"))
	fmt.Fprintln(os.Stderr, "A tool to help with conventional commits and more.")
	fmt.Fprintln(os.Stderr, "\n"+headerStyle.Render("USAGE"))
	fmt.Fprintln(os.Stderr, "  ch <command> [arguments]")
	fmt.Fprintln(os.Stderr, "\n"+headerStyle.Render("COMMANDS"))
	fmt.Fprintln(os.Stderr, "  "+commandStyle.Render("commit")+":   Run the interactive commit builder.")
	fmt.Fprintln(os.Stderr, "  "+commandStyle.Render("validate")+" <file>: Validate a commit message file.")
}
