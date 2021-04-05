package solutions

func CountIf(f func(string) bool, a []string) int {
	counter := 0
	for _, el := range a {
		if f(el) {
			counter++
		}
	}

	return counter
}
