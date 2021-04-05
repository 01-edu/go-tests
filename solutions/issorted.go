package solutions

func IsSorted(f func(int, int) int, a []int) bool {
	ascendingOrdered := true
	descendingOrdered := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			ascendingOrdered = false
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			descendingOrdered = false
		}
	}

	return ascendingOrdered || descendingOrdered
}
