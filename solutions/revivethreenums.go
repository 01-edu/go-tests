package solutions

func max(ints ...int) int {
	if len(ints) == 0 {
		return 0
	}
	max := ints[0]
	for _, i := range ints[1:] {
		if i > max {
			max = i
		}
	}
	return max
}

func ReviveThreeNums(a, b, c, d int) int {
	maxi := -111
	if a != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-a)
	}
	if b != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-b)
	}
	if c != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-c)
	}
	if d != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-d)
	}
	return maxi
}
