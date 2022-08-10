package solutions

func EvenLength(strings []string) []string {
	var newSlice []string
	for _, str := range strings {
		if len(str)%2 == 0 {
			newSlice = append(newSlice, str)
		}
	}
	return newSlice
}
