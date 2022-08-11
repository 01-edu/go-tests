package solutions

import "strings"

func PrintAndMiss(arg string, loop int) string {
	if arg == "" || loop < 0 {
		return "Invalid Output\n"
	}
	if loop == 0 || loop > len(arg) {
		return arg + "\n"
	}
	var _str strings.Builder
	for i := 0; i < len(arg); i++ {
		if i != 0 && i%loop == 0 {
			i += loop
			if i > len(arg)-1 {
				break
			}
		}
		if i != len(arg) {
			_str.WriteRune(rune(arg[i]))
		}
	}
	_str.WriteRune(rune('\n'))
	return (_str.String())
}
