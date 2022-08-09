package solutions

func SliceAdd(slice []int, num int) []int {
	if len(slice) == 0 {
		return []int{num}
	}
	slice = append(slice, num)
	return slice
}
