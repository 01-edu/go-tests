package solutions

func PopInt(ints []int) []int {
	if len(ints) == 0 {
		return ints
	}
	return ints[:len(ints)-1]
}
