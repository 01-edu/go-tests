package solutions

import "fmt"

func BuzZinga(n int32) {
	var i int32 = 0
	if n == 0 {
		fmt.Println("Buzzinga")
	}
	if n < 0 {
		fmt.Println("*")
	}
	for i = 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Buzzinga")
		} else if i%3 == 0 {
			fmt.Println("Buz")
		} else if i%5 == 0 {
			fmt.Println("zinga")
		} else {
			fmt.Println("*")
		}
	}
}
