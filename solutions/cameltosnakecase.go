package solutions

func containOnlyAlphabe(str string) bool {
	for _, c := range str {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}

func countUpperCaseAlphabet(str string) int {
	count := 0
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			count++
		}
	}
	return count
}

func CamelToSnakeCase(s string) string {
	result := ""
	if len(s) == 0 || s[0] < 'a' || s[0] > 'z' || !containOnlyAlphabe(s) || countUpperCaseAlphabet(s) != 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += "_"
			result += string(s[i])
		} else {
			result += string(s[i])
		}
	}
	return result
}
