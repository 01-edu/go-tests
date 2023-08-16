package main

func StringToIntSlice(str string) []int {
	runes := []rune(str)

	var result []int

	for _, r := range runes {
		result = append(result, int(r))
	}
	return result
}
