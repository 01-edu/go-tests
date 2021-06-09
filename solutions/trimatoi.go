package solutions

import (
	"strconv"
)

func TrimAtoi(s string) int {
	var s2 string
	for _, r := range s {
		if (r >= '0' && r <= '9') || (r == '-' && s2 == "") {
			s2 += string(r)
		}
	}
	a, _ := strconv.Atoi(s2)
	return a
}
