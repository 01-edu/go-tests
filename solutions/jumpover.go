package solutions

import "strings"

func JumpOver(arg string) string {
	if arg == "" || len(arg) < 3 {
		return "\n"
	}
	var str strings.Builder
	for i := 0; i < len(arg); i++ {
		if i != 0 {
			j := i + 1
			if j%3 == 0 {
				str.WriteRune(rune(arg[i]))
			}
		}
	}
	str.WriteRune(rune('\n'))
	return (str.String())
}
