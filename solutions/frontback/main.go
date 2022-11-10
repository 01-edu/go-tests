package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 && len(os.Args[1]) == 1 {
		char := os.Args[1][0]
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if char == 'z' || char == 'Z' {
				z01.PrintRune(rune(char) - 1)
				z01.PrintRune(rune(char))
			} else if char == 'a' || char == 'A' {
				z01.PrintRune(rune(char))
				z01.PrintRune(rune(char) + 1)
			} else {
				z01.PrintRune(rune(rune(char)) - 1)
				z01.PrintRune(rune(char))
				z01.PrintRune(rune(char) + 1)
			}
		}
	}
	z01.PrintRune('\n')
}
