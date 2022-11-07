package solutions

func PrintIfNot(arg string) string {
	if len(arg) >= 3 {
		return "Invalid Input\n"
	} else {
		return "G\n"
	}
}
