package solutions

func Numofdigits(n int) int {
	var count int = 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}
