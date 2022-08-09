package solutions

func SliceRemove(slice []int, num int) []int {
	newSlice := []int{}
	for _, v := range slice {
		if v != num {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
