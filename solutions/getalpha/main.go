package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		i, err := strconv.Atoi(os.Args[1])
		if err == nil && i >= 0 && i <= 127 {
			z01.PrintRune(rune(i))
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}
