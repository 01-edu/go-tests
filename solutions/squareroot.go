package solutions

func SquareRoot(number int) int {
	if number == 1{
		return 1
	}
	if number >= 0 {
		sqrt := number / 2
		temp := 0
		for sqrt != temp {
			temp = sqrt
			sqrt = (number/temp + temp) / 2
		}
		return sqrt
	}
	return -1
}
