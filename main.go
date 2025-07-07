package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// --- STYLES ---
// We'll define our styles using Lip Gloss.

var (
	// Style for the main container/box
	guidelineBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")). // Purple
				Padding(1, 2).
				MarginBottom(1)

	// Style for error messages
	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")). // Red
			Bold(true)

	// Style for success messages
	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")). // Green
			Bold(true)

	// Style for code snippets and examples
	codeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")). // Gray
			Background(lipgloss.Color("237"))

	// Style for titles within the guideline box
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")). // Blue
			Bold(true).
			Underline(true)
)

func main() {
	// Define the allowed commit types
	allowedTypes := map[string]bool{
		"feat":     true,
		"fix":      true,
		"build":    true,
		"chore":    true,
		"ci":       true,
		"docs":     true,
		"style":    true,
		"refactor": true,
		"perf":     true,
		"test":     true,
		"revert":   true,
	}

	// Check for the correct number of arguments
	if len(os.Args) < 2 {
		fmt.Println(errorStyle.Render("[ERROR] Usage:"), os.Args[0], "<path_to_commit_message_file>")
		os.Exit(1)
	}
	msgFilePath := os.Args[1]

	// Read the commit message from the file
	messageBytes, err := os.ReadFile(msgFilePath)
	if err != nil {
		fmt.Println(errorStyle.Render("[ERROR]"), "Failed to read commit message file:", msgFilePath)
		os.Exit(1)
	}
	message := string(messageBytes)

	// Ignore comments and get the first line
	firstLine := strings.TrimSpace(strings.Split(message, "\n")[0])

	if firstLine == "" {
		fmt.Println(errorStyle.Render("[ERROR] Commit message cannot be empty."))
		printGuidelines()
		os.Exit(1)
	}

	// Regex to validate the Conventional Commits format
	re := regexp.MustCompile(`^(?P<type>\w+)(?P<scope>\(\w+\))?(?P<breaking>!)?:(?P<subject>.+)$`)
	matches := re.FindStringSubmatch(firstLine)
	if matches == nil {
		fmt.Println(errorStyle.Render("[ERROR] Invalid commit message format."))
		printGuidelines()
		os.Exit(1)
	}

	// Extract parts from the regex match
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" && i < len(matches) {
			result[name] = matches[i]
		}
	}

	// Validate the commit type
	commitType := result["type"]
	if !allowedTypes[commitType] {
		fmt.Println(errorStyle.Render("[ERROR] Invalid commit type:"), fmt.Sprintf("'%s'", commitType))
		printGuidelines()
		os.Exit(1)
	}

	// Validate the subject
	subject := result["subject"]
	if !strings.HasPrefix(subject, " ") {
		fmt.Println(errorStyle.Render("[ERROR] Subject must have a leading space after the colon."))
		printGuidelines()
		os.Exit(1)
	}
	if strings.TrimSpace(subject) == "" {
		fmt.Println(errorStyle.Render("[ERROR] Subject cannot be empty."))
		printGuidelines()
		os.Exit(1)
	}

	// If all checks pass
	fmt.Println(successStyle.Render("[SUCCESS]"), "Commit message is valid.")
	os.Exit(0)
}

// printGuidelines prints the beautifully styled help text.
func printGuidelines() {
	var sb strings.Builder

	sb.WriteString(titleStyle.Render("Conventional Commits Guidelines") + "\n\n")
	sb.WriteString("A valid commit message follows this format:\n")
	sb.WriteString(codeStyle.Render("<type>[optional scope]: <description>") + "\n\n")
	sb.WriteString(titleStyle.Render("Examples") + "\n")
	sb.WriteString("  " + codeStyle.Render("feat(api): implement user authentication") + "\n")
	sb.WriteString("  " + codeStyle.Render("fix: correct a bug in the login flow") + "\n\n")
	sb.WriteString("For more details, see: https://www.conventionalcommits.org")

	fmt.Println(guidelineBoxStyle.Render(sb.String()))
}
