package solutions

func IsSquare(number int, square int) bool {
	if number < 0 || square < 0 {
		return false
	}
	return number*number == square
}
