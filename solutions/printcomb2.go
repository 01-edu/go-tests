package solutions

import "fmt"

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			fmt.Printf("%.2d %.2d", a, b)
			if a != 98 || b != 99 {
				fmt.Print(", ")
			}
		}
	}
	fmt.Println()
}
