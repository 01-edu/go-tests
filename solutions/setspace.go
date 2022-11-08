package solutions

func SetSpace(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' && i != 0 {
			res += " " + string(str[i])
		} else if str[i] >= 'a' && str[i] <= 'z' && i != 0 || i == 0 && str[i] >= 'A' && str[i] <= 'Z' {
			res += string(str[i])
		} else {
			return "Error"
		}
	}
	return res
}
