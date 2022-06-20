package solutions

func SquareRoot(number int) int {
	if number == 1 {
		return 1
	}
	if number >= 0 {
		var sqrt float32 = float32(number / 2)
		var temp float32 = 0.0
		for sqrt != temp {
			temp = sqrt
			sqrt = float32((float32(number)/temp + temp) / 2)
		}
		return int(sqrt)
	}
	return -1
}
