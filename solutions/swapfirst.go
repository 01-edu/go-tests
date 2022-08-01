package solutions

func SwapFirst(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	slice[0], slice[1] = slice[1], slice[0]
	return slice
}
