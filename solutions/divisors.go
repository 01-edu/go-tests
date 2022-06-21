package solutions

func Divisors(n int) int {
	var count int = 0
	if n > 0 {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				count++
			}
		}
	}
	return count
}
