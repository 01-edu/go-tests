package solutions

func AlphaPosition(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 1)
	}
	return -1
}
