package solutions

import (
	"strconv"
)

func Sum(a, b string) int {
	numA, _ := strconv.Atoi(a)
	numB, _ := strconv.Atoi(b)
	if numA > 9 || numA < -9 || numB > 9 || numB < -9 {
		return 0
	}
	return numA + numB
}
