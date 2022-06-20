package solutions

func IsMultiple(num int) bool {
	if num > 0 && (num%3 == 0 || num%7 == 0) {
		return true
	}
	return false
}
