package solutions

func SwapFirst(slice []int) []int {
	dst := make([]int, len(slice))
	copy(dst, slice)

	if len(dst) < 2 {
		return dst
	}
	dst[0], dst[1] = dst[1], dst[0]
	return dst
}
