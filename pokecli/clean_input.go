package pokecli

import (
	"strings"
)

func CleanInput(text string) []string {
	text = strings.Trim(text, " ")
	if len(text) == 0 {
		return make([]string, 0)
	}

	parts := strings.Fields(text)
	for idx := range parts {
		parts[idx] = strings.ToLower(parts[idx])
	}
	return parts 
}