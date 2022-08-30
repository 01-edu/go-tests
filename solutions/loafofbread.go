package solutions

import "strings"

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var _str strings.Builder
	j := 0
	for i := 0; i < len(str); i++ {
		if j < 5 && rune(str[i]) == ' ' {
			continue
		}
		if j == 5 {
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}
			if i == len(str)-1 {
				break
			}
			_str.WriteRune(' ')
			j = 0
			continue
		}
		_str.WriteRune(rune(str[i]))
		j++
	}
	_str.WriteRune(rune('\n'))
	return (_str.String())
}
