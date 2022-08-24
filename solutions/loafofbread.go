package solutions

import "strings"

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	s := strings.ReplaceAll(str, " ", "")
	var _str strings.Builder
	j := 0
	for i := 0; i < len(s); i++ {
		if j == 5 {
			_str.WriteRune(rune(' '))
			j = 0
		} else {
			_str.WriteRune(rune(s[i]))
			j++
		}
	}
	_str.WriteRune(rune('\n'))
	return (_str.String())
}
