package solutions

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	acc := a[0]
	for i := 1; i < len(a); i++ {
		acc = f(acc, a[i])
	}
	fmt.Println(acc)
}
