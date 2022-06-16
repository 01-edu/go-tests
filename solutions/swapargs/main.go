package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}
	for i := 0; i < len(os.Args[2]); i++ {
		z01.PrintRune(rune(os.Args[2][i]))
	}
	for i := 0; i < len(os.Args[1]); i++ {
		z01.PrintRune(rune(os.Args[1][i]))
	}
	z01.PrintRune('\n')
}
