package solutions

func AddFront(str string, slice []string) []string {
	return append([]string{str}, slice...)
}
