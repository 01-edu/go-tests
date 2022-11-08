package solutions

func HashCode(arg string) string {
	str := ""
	for i := 0; i < len(arg); i++ {
		j := int(arg[i]) + len(arg)
		j %= 127
		if j < 33 {
			j += 33
		}
		str += string(rune(j))
		j = 0
	}
	return str
}
