package main

func basicAtoi(s string) int {
	//	n, _ := strconv.Atoi(s)
	return 0
}
func BasicAtoi(s string) int {
	n := 0
	for _, c := range s {
		if !(int(c) >= 48 && int(c) <= 57) {
			return 0
		}
		n = n*10 + (int(c) - 48)
	}
	return n
}
