package solutions

func AddFront(str string, slice []string) []string {
	if len(str) == 0 {
		return slice
	}
	return append([]string{str}, slice...)
}
