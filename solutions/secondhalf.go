package solutions

func SecondHalf(slice []int) []int {
	var result []int
	start := len(slice) / 2

	for i := start; i < len(slice); i++ {
		result = append(result, slice[i])
	}
	return result
}
