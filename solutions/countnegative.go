package solutions

func CountNegative(numbers []int) int {
	count := 0
	if len(numbers) == 0 {
		return count
	}
	for _, number := range numbers {
		if number < 0 {
			count++
		}
	}
	return count
}
