package main

// allowedTypes defines the valid commit types.
var allowedTypes = []string{
	"feat", "fix", "build", "chore", "ci", "docs", "style",
	"refactor", "perf", "test", "revert",
}

// isAllowedType checks if a given type is in the list of allowed types.
func isAllowedType(commitType string, types []string) bool {
	for _, t := range types {
		if t == commitType {
			return true
		}
	}
	return false
}
