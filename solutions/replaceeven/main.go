package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for argsIndex, arg := range os.Args[1:] {
		for i := 0; i < len(arg); i++ {
			if i%2 == 0 {
				z01.PrintRune('2')
			} else {
				z01.PrintRune(rune(arg[i]))
			}
		}
		if argsIndex < len(os.Args[1:])-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
