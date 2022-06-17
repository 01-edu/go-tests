package solutions

func BinaryAddition(a, b int) []int {
	var binaryA []int
	var binaryB []int
	if a < 0 || b < 0 {
		return nil
	}
	for a > 0 {
		binaryA = append(binaryA, a%2)
		a = a / 2
	}
	for b > 0 {
		binaryB = append(binaryB, b%2)
		b = b / 2
	}
	var sum []int
	carry := 0
	i := 0
	for i < len(binaryA) || i < len(binaryB) {
		if i < len(binaryA) {
			carry += binaryA[i]
		}
		if i < len(binaryB) {
			carry += binaryB[i]
		}
		sum = append(sum, carry%2)
		carry = carry / 2
		i++
	}
	if carry > 0 {
		sum = append(sum, carry)
	} else if carry == 0 {
		sum = append(sum, 0)
	}
	for i, j := 0, len(sum)-1; i < j; i, j = i+1, j-1 {
		sum[i], sum[j] = sum[j], sum[i]
	}
	return sum
}
