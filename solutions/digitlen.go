package solutions

func DigitLen(num, base int) int {
	if num < 0 {
		num = -num
	}
	if base < 2 || base > 36 {
		return -1
	}
	res := 0
	for num > 0 {
		num /= base
		res++
	}
	return res
}
