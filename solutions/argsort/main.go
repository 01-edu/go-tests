package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	if len(argument) == 1 && len(argument[0]) >= 1 {
		for i := 0; i < len(argument[0]); i++ {
			j := i + 1
			if j < len(argument[0]) {
				if argument[0][i] > argument[0][j] {
					z01.PrintRune('F')
					z01.PrintRune('\n')
					return
				}
			}
		}
		z01.PrintRune('T')
		z01.PrintRune('\n')
	}
}
