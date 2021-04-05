package solutions

func TrimAtoi(s string) int {
	runes := []rune(s)
	for i := range s {
		runes[i] = rune(s[i])
	}
	var numbers []rune
	for _, r := range runes {
		if (r >= '0' && r <= '9') || (len(numbers) == 0 && (r) == '-' || r == '+') {
			numbers = append(numbers, r)
		}
	}
	if len(numbers) == 0 || (len(numbers) == 1 && (numbers[0] == '-' || numbers[0] == '+')) {
		return 0
	}

	res, i, sign := 0, 0, 1

	if numbers[0] == '-' {
		sign = -1
		i++
	} else if numbers[0] == '+' {
		i++
	}
	for ; i < len(numbers); i++ {
		res = res*10 + int(numbers[i]) - '0'
	}

	return sign * res
}
