package solutions

func SwapLast(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	slice[len(slice)-1], slice[len(slice)-2] = slice[len(slice)-2], slice[len(slice)-1]
	return slice
}
