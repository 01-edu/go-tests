package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/z01"
	// "github.com/01-edu/go-tests/lib/random"
)

// func solutions(n int, n2 int) string {
// 	solutions := map[[2]int]string{
// 		[2]int{1, 10}:   "1 2 3 4 5 6 7 8 9 9\n",
// 		[2]int{10, 1}:   "10 9 8 7 6 5 4 3 2 1\n",
// 		[2]int{1, 1}:    "1\n",
// 		[2]int{10, 10}:  "\n",
// 		[2]int{0, 9}:    "0 1 2 3 4 5 6 7 8 9\n",
// 		[2]int{-1, -10}: "\n",
// 		[2]int{-10, -1}: "\n",
// 		[2]int{-1, 9}:   "0 1 2 3 4 5 6 7 8 9\n",
// 	}
// 	return solutions[[2]int{n, n2}]
// }


func printRange(n, m int) {
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


func main() {
	args := [][]int{
		{1, 10},
		{10, 1},
		{1, 1},
		{10, 10},
		{0, 9},
		{-1, -10},
		{-10, -1},
		{-1, 9},
	}
	for _, arg := range args {
	challenge.Function("printRange", student.PrintRange,printRange, arg[0], arg[1])
	}
}
