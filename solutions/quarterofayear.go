package solutions

func QuarterOfAYear(month int) int {
	if month < 1 || month > 12 {
		return -1
	}
	return (month + 2) / 3
}
