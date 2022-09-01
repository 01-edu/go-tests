package solutions

func ReverseMenuIndex(menu []string) []string {
	menuLen := len(menu)
	output := make([]string, menuLen)

	for i, n := range menu {
		j := menuLen - i - 1

		output[j] = n
	}
	return output
}
