package solutions

func RectPerimeter(w, h int) int {
	if (w < 0) || (h < 0) {
		return -1
	}

	return 2 * (w + h)
}
