package solutions

func containOnlyAlphabet(str string) bool {
	for _, c := range str {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}

func isUpper(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func CamelToSnakeCase(s string) string {
	result := ""
	if len(s) == 0 || !containOnlyAlphabet(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i != 0 && isUpper(rune(s[i])) && i+1 < len(s) && !isUpper(rune(s[i+1])) {
			result += "_"
			result += string(s[i])
		} else if !isUpper(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			result += string(s[i])
		} else {
			return s
		}
	}
	return result
}
