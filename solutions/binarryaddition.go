package solutions

func BinaryAddition(a, b int) []int {
	var result []int
	var carry int = 0
	for a > 0 || b > 0 {
		var sum int
		if a > 0 {
			sum = a % 2
			a = a / 2
		}
		if b > 0 {
			sum += b % 2
			b = b / 2
		}
		sum += carry
		carry = sum / 2
		result = append(result, sum%2)
	}
	if carry > 0 {
		result = append(result, carry)
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
