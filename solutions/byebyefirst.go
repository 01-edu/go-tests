package solutions

func ByeByeFirst(strings []string) []string {
	if len(strings) <= 1 {
		return []string{}
	}
	return strings[1:]
}
