package solutions

func PrintAndMiss(arg string, num int) string {
	if arg == "" || num < 0 {
		return "Invalid Output\n"
	}
	if num == 0 || num > len(arg) {
		return arg + "\n"
	}
	_str := ""
	for i := 0; i < len(arg); i++ {
		if i != 0 && i%num == 0 {
			i += num
			if i > len(arg)-1 {
				break
			}
		}
		if i != len(arg) {
			_str += string(rune(arg[i]))
		}
	}
	return (_str + "\n")
}
