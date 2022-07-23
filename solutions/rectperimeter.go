package solutions

func RectPerimeter(w, l int) int {
	if (w < 0) || (l < 0) {
		return -1
	}

	return 2 * (w + l)
}
