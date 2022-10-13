package solutions

import (
	"math"
	"strconv"
)

func NotDecimal(dec string) string {
	j := -1
	n := 0
	if len(dec) == 0 {
		return "\n"
	}
	for i := 0; i < len(dec); i++ {
		if j == -1 && dec[i] == '.' {
			j++
		} else if j == 0 {
			n++
		}
	}
	s, err := strconv.ParseFloat(dec, 64)
	if err == nil {
		return strconv.Itoa(int(s*math.Pow(10, float64(n)))) + "\n"
	}
	return (dec + "\n")
}
