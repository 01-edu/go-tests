package solutions

func Fibonacci(i int) int {
	if i < 0 {
		return -1
	}
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}
