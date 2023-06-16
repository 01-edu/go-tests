package solutions

import "strings"

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return (str)
	} else {
		var res strings.Builder
		i := 0
		for i = 0; i < int(len(str)/2); i++ {
			res.WriteRune(rune(str[i]))
		}
		return res.String()
	}
}
