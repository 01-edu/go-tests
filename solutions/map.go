package solutions

func Map(f func(int) bool, a []int) (result []bool) {
	for _, el := range a {
		result = append(result, f(el))
	}
	return
}
