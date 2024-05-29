package solutions

import "strings"

func FirstWord(s string) string {
	words := strings.Fields(s)
	res := "\n"
	if len(words) > 0 {
		res = words[0] + res
	}
	return res
}
