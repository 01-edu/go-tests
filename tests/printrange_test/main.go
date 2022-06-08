package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	// "github.com/01-edu/go-tests/lib/random"
)

func solutions(n int, n2 int) string {
	solutions := map[[2]int]string{
		[2]int{1, 10}:   "1 2 3 4 5 6 7 8 9 9\n",
		[2]int{10, 1}:   "10 9 8 7 6 5 4 3 2 1\n",
		[2]int{1, 1}:    "1\n",
		[2]int{10, 10}:  "\n",
		[2]int{0, 9}:    "0 1 2 3 4 5 6 7 8 9\n",
		[2]int{-1, -10}: "\n",
		[2]int{-10, -1}: "\n",
		[2]int{-1, 9}:   "0 1 2 3 4 5 6 7 8 9\n",
	}
	return solutions[[2]int{n, n2}]
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
	challenge.Function("printRange", student.printRange,solutions, args)
}
