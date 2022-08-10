package solutions

func Delete(ints []int, position int) []int {
	if position > len(ints)-1 {
		return ints
	}
	result := []int{}
	for i, integer := range ints {
		if i == position-1 {
			continue
		}
		result = append(result, integer)
	}
	return result
}
