package solutions

func CountNegative(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if number < 0 {
			count++
		}
	}
	return count
}
