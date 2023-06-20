package solutions

func SwapLast(slice []int) []int {
	dst := make([]int, len(slice))
	copy(dst, slice)

	if len(dst) < 2 {
		return dst
	}
	dst[len(dst)-1], dst[len(dst)-2] = dst[len(dst)-2], dst[len(dst)-1]
	return dst
}
