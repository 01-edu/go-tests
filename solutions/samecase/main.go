package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[1:]
	if len(argument) == 2 && len(argument[0]) == 1 && len(argument[1]) == 1 {
		if ((argument[0][0] >= 'a' && argument[0][0] <= 'z') || (argument[0][0] >= 'A' && argument[0][0] <= 'Z')) &&
			(argument[1][0] >= 'a' && argument[1][0] <= 'z') || (argument[1][0] >= 'A' && argument[1][0] <= 'Z') {
			if (argument[0][0] >= 'A' && argument[0][0] <= 'Z') && (argument[1][0] >= 'A' && argument[1][0] <= 'Z') ||
				(argument[0][0] >= 'a' && argument[0][0] <= 'z') && (argument[1][0] >= 'a' && argument[1][0] <= 'z') {
				z01.PrintRune('1')
			} else {
				z01.PrintRune('0')
			}
		} else {
			z01.PrintRune('-')
			z01.PrintRune('1')
		}
	}
	z01.PrintRune('\n')
}
