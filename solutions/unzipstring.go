package solutions

import "strconv"

func check(a rune) bool {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z'){
		return true
	}
	return false
}

func Unzipstring(str string) string {
	var result string
	if len(str) == 0 {
		return "Invalid output"
	}
	num, err := strconv.Atoi(string(str[0]))
	if err != nil {
		return "Invalid output"
	}
	for i := 0; i < num; i++ {
		if check(rune(str[1])) == true {
			result += string(str[1])
		} else {
			return "Invalid output"
		}
	}
	if len(str) > 2 {
		if Unzipstring(string(str[2:])) != "Invalid output" {
			result += Unzipstring(string(str[2:]))
		} else {
			return "Invalid output"
		}
	}
	return result
}
