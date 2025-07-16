package main

import "github.com/charmbracelet/lipgloss"

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
			Background(lipgloss.Color("237")).
			Padding(0, 1)

	// Style for titles within the guideline box
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("12")). // Blue
			Bold(true).
			Underline(true)

	// Style for faint/help text
	faintStyle = lipgloss.NewStyle().
			Faint(true)

	// Style for the help box
	helpBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("241")).
			Padding(1, 2).
			MarginTop(1)
)
