package solutions

func RemoveOdd(str string) string {
	var result string
	for  i:=0; i<len(str); i++ {
		if (i+1)%2 != 0 || str[i] == ' ' {
			result += string(str[i])
		}
	}
	return result
}
