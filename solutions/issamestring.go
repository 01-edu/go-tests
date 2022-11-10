package solutions

import "unicode"

func getAlphabetIndex(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A')
	}
	return -1
}

func IsSameString(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	i := 0
	for i < len(s1) && i < len(s2) {
		if (unicode.IsLetter(rune(s1[i])) && unicode.IsLetter(rune(s2[i]))) &&
			(getAlphabetIndex(s1[i]) != getAlphabetIndex(s2[i])) {
			return false
		} else if !(unicode.IsLetter(rune(s1[i])) && unicode.IsLetter(rune(s2[i]))) && s1[i] != s2[i] {
			return false
		}
		i++
	}
	if i < len(s1) || i < len(s2) {
		return false
	}
	return true
}
