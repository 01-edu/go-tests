package solutions

func PrintIf(arg string) string {
	if len(arg) != 0 && len(arg) < 3 {
		return "Invalid Input\n"
	}
	return "G\n"
}
