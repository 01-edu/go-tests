package solutions

func PriorPrime(x int) (ans int) {
	var ok bool
	for i := 2; i < x; i++ {
		ok = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				ok = false
			}
		}
		if ok {
			ans += i
		}
	}
	return
}
