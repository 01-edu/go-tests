package solutions

func Oddlength(strings []string) []string {
	var newSlice []string
	for _, str := range strings {
		if len(str)%2 == 1 {
			newSlice = append(newSlice, str)
		}
	}
	return newSlice
}
