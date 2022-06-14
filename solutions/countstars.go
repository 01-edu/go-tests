package solutions

import "strconv"

func CountStars(num int) string {
	res := ""
	star := " star..."
	if num <= 0 {
		res = "No stars"
	}
	for i := 1; i <= num; i++ {
		res = res + strconv.Itoa(i) + star
	}
	return res
}
