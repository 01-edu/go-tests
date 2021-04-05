package solutions

func Unmatch(elems []int) int {
	var quant int
	for _, el := range elems {
		quant = 0
		for _, v := range elems {
			if v == el {
				quant++
			}
		}
		if quant%2 != 0 {
			return el
		}
	}
	return -1
}
