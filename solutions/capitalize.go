package solutions

import (
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	r := []rune(strings.ToLower(s))

	if unicode.Is(unicode.Latin, r[0]) {
		r[0] = unicode.ToUpper(r[0])
	}

	for i := 1; i < len(r); i++ {
		if !unicode.IsDigit(r[i-1]) && !unicode.Is(unicode.Latin, r[i-1]) {
			r[i] = unicode.ToUpper(r[i])
		}
	}
	return string(r)
}
