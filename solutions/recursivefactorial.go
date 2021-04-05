package solutions

func RecursiveFactorial(nb int) int {
	limit := 20
	if nb < 0 || nb > limit {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
