package solutions

import "github.com/01-edu/z01"

func PrintRange(n, m int) {
	if n < 0 && m < 0 || n > 9 && m > 9 {
		z01.PrintRune('\n')
		return
	}
	if n < 0 {
		n = 0
	} else if n > 9 {
		n = 9
	}
	if m < 0 {
		m = 0
	} else if m > 9 {
		m = 9
	}
	if n > m {
		for i := n; i >= m; i-- {
			z01.PrintRune(rune(i) + '0')
			if i != m {
				z01.PrintRune(' ')
			}
		}
	} else {
		for i := n; i <= m; i++ {
			z01.PrintRune(rune(i) + '0')
			if i != m {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
