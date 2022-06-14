package solutions

func SumArray(array []int) int {
	var sum int = 0
	for _, v := range array {
		sum += v
	}
	return sum
}
