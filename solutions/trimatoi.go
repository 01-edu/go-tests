package solutions

import (
	"strconv"
	"unicode"
)

func recursion(s string, i int) string {
	result := ""
	if i < 0 {
		return ""
	}
	if unicode.IsDigit(rune(s[i])) {
		result += string(s[i])
	}
	if s[i] == '-' && recursion(s[:i], i-1) == "" {
		return "-"
	}
	return recursion(s[:i], i-1) + result
}

func TrimAtoi(s string) int {
	if a, err := strconv.Atoi(recursion(s, len(s)-1)); err == nil {
		return a
	}
	return 0
}
