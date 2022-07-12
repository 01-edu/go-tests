package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var vowels []rune

	arguments := os.Args[1:]

	for _, i := range arguments {
		for j := 0; j < len(i); j++ {
			if IsVowel(i[j]) {
				vowels = append(vowels, rune(i[j]))
			}
		}
	}

	counter := len(vowels) - 1

	for s, i := range arguments {
		for j := 0; j < len(i); j++ {
			if IsVowel(i[j]) {
				z01.PrintRune(rune(vowels[counter]))
				counter--
			} else {
				z01.PrintRune(rune(i[j]))
			}
		}
		if s != len(arguments)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func IsVowel(s byte) bool {
	if s == 65 || s == 69 || s == 73 || s == 79 || s == 85 || s == 97 || s == 101 || s == 105 || s == 111 || s == 117 {
		return true
	}
	return false
}
