package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{},
		{1},
		{1, 2},
		{1, 2, 3},
	}
	for _, arg := range args {
		challenge.Function("SwapLast", student.SwapLast, solutions.SwapLast, arg)
	}
}
