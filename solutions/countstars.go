package solutions

import "strconv"

func CountStars(num int) string {
	res := ""
	star := " star"
	stars := " stars"
	dots := "..."
	if num <= 0 {
		return "No stars"
	}
	if num == 1 {
		return strconv.Itoa(1) + star
	}
	for i := 1; i <= num; i++ {
		if i == 1 {
			res = res + strconv.Itoa(i) + star + dots
		} else if i != num {
			res = res + strconv.Itoa(i) + stars + dots
		} else {
			res = res + strconv.Itoa(i) + stars
		}
	}
	return res
}
