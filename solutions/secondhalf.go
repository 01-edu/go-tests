package solutions

func SecondHalf(slice []int) []int {
	var result []int
	size := len(slice)

	if size%2 == 0 {
		size /= 2
	} else {
		size /= 2
	}
	for i := size; i < len(slice); i++ {
		result = append(result, slice[i])
	}
	return result
}
