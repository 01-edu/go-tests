package solutions

func BeZero(slice []int) []int {
	dst := make([]int, len(slice))
	copy(dst, slice)

	for i := 0; i < len(dst); i++ {
		dst[i] = 0
	}
	return dst
}
