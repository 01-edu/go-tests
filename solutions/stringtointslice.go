package solutions

func StringToIntSlice(str string) []int {
	runes := []rune(str)

	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	return result
}
