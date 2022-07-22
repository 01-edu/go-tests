package solutions

import (
	"fmt"
	"strconv"
)

func StrisNegative(str string) {
	if len(str) > 0 {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("!")
		} else {
			if val < 0 {
				fmt.Println("N")
			} else if val == 0 {
				fmt.Println("0")
			} else {
				fmt.Println("P")
			}
		}
	} else {
		fmt.Println("!")
	}
}
