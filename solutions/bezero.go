package solutions

func BeZero(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] = 0
	}
	return slice
}
