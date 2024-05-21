package solutions

import (
	"strings"
)

func LastWord(s string) string {
	words := strings.Fields(s)
	if len(words) > 0 {
		return words[len(words)-1] + "\n"
	}
	return "\n"
}
