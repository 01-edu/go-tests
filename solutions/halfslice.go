package solutions

func HalfSlice(slice []int) []int {
	var result []int
	size := len(slice) / 2

	if len(slice)%2 != 0 {
		size = len(slice)/2 + 1
	}
	for i := 0; i < size; i++ {
		result = append(result, slice[i])
	}
	return result
}
