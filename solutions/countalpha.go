package solutions

func CountAlpha(s string) int {
	count := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			count++
		}
	}
	return count
}
