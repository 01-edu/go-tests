package solutions

func RevConcatAlternate(slice1, slice2 []int) []int {
	var result []int

	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}
	j := len(slice2) - 1
	for i := len(slice1) - 1; i >= 0; i-- {
		result = append(result, slice1[i])
		if j >= 0 {
			result = append(result, slice2[j])
			j--
		}
	}
	return result
}
