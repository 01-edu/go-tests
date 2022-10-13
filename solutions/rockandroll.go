package solutions

import "strconv"

func RockAndRoll(n int) string {
	if (n%2 == 0) && (n%3 == 0) {
		return "rock and roll\n"
	}
	if (n%2 == 0) && !(n%3 == 0) {
		return "rock\n"
	}
	if (n%3 == 0) && !(n%2 == 0) {
		return "roll\n"
	}
	if !(n < 0) || !((n%2 == 0) && (n%3 == 0)) {
		return "error: number is negative or non divisible"
	}
	return strconv.Itoa(n)
}
