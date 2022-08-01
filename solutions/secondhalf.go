package solutions

func SecondHalf(slice []int) []int {
	var result []int
	size := len(slice)

	if len(slice)%2 == 0 {
		size = len(slice) / 2
	} else {
		size = len(slice) / 2
	}
	for i := size; i < len(slice); i++ {
		result = append(result, slice[i])
	}
	return result
}
