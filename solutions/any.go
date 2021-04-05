package solutions

func Any(f func(string) bool, a []string) bool {
	for _, el := range a {
		if f(el) {
			return true
		}
	}
	return false
}
