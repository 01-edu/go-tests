package solutions

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func RemoveDuplicate(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	newSlice := []int{}
	for _, v := range slice {
		if !contains(newSlice, v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
