package solutions

import "unicode"

func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}
