package solutions

import "github.com/01-edu/z01"

func printInteger(a int) {
	secondDigits := (a % 10)
	firstDigits := (a / 10)

	z01.PrintRune(rune(firstDigits) + 48)
	z01.PrintRune(rune(secondDigits) + 48)
}

func DescendComb() {
	for i := 99; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			printInteger(i)
			z01.PrintRune(' ')
			printInteger(j)

			if i == 1 && j == 0 {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}
