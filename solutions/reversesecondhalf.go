package solutions

func ReverseSecondHalf(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	} else if len(str) == 1 {
		return (str + "\n")
	} else {
		res := ""
		i := 0
		for i = len(str) - 1; i >= int(len(str)/2); i-- {
			res += string(str[i])
		}
		res += "\n"
		return res
	}
}
