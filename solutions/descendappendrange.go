package solutions

func DescendAppendRange(max, min int) []int {
	answer := []int{}
	if max <= min {
		return []int{}
	} else {
		for i := max ; i > min; i-- {
			answer = append(answer, i)
		}
	}
	return answer
}
