package solutions

import "strings"

func PrintFirstHalf(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	} else if len(str) == 1 {
		return (str + "\n")
	} else {
		var res strings.Builder
		i := 0
		for i = 0; i < int(len(str)/2); i++ {
			res.WriteRune(rune(str[i]))
		}
		res.WriteRune(rune('\n'))
		return res.String()
	}
}
