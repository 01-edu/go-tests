package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}
	str := ""
	for i := 0; i < len(args); i++ {
		for _, v := range args[i] {
			str += string(v)
		}
	}
	if len(str) > 0 {
		for _, v := range str {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
