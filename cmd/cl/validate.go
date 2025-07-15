package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func runValidationMode(msgFilePath string) error {
	messageBytes, err := os.ReadFile(msgFilePath)
	if err != nil {
		return fmt.Errorf("failed to read commit message file: %s", msgFilePath)
	}
	message := string(messageBytes)

	firstLine := strings.TrimSpace(strings.Split(message, "\n")[0])

	if firstLine == "" {
		fmt.Println(errorStyle.Render("[ERROR] Commit message cannot be empty."))
		printGuidelines()
		return fmt.Errorf("empty commit message")
	}

	// Regex to validate the Conventional Commits format
	re := regexp.MustCompile(`^(?P<type>\w+)(?P<scope>\(\w+\))?(?P<breaking>!)?:(?P<subject>.+)$`)
	matches := re.FindStringSubmatch(firstLine)
	if matches == nil {
		fmt.Println(errorStyle.Render("[ERROR] Invalid commit message format."))
		printGuidelines()
		return fmt.Errorf("invalid format")
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" && i < len(matches) {
			result[name] = matches[i]
		}
	}

	commitType := result["type"]
	if !isAllowedType(commitType, allowedTypes) {
		fmt.Println(errorStyle.Render("[ERROR] Invalid commit type:"), fmt.Sprintf("'%s'", commitType))
		fmt.Println("\nAllowed types are:")
		fmt.Println("  " + codeStyle.Render(strings.Join(allowedTypes, ", ")))
		printGuidelines()
		return fmt.Errorf("invalid type")
	}

	subject := result["subject"]
	if !strings.HasPrefix(subject, " ") {
		fmt.Println(errorStyle.Render("[ERROR] Subject must have a leading space after the colon."))
		printGuidelines()
		return fmt.Errorf("invalid subject format")
	}
	if strings.TrimSpace(subject) == "" {
		fmt.Println(errorStyle.Render("[ERROR] Subject cannot be empty."))
		printGuidelines()
		return fmt.Errorf("empty subject")
	}

	return nil
}

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
