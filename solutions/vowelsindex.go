package solutions

func VowelsIndex(str string) []int {
	var res []int
	k := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
			j := i
			if j < len(str) {
				res = append(res, j)
				k++
			}
		}
	}
	return res
}
