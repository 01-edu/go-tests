package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{1, 2, 3},
		{1, 1, 2},
		{10, 20, 30},
		{20, 10, 30},
		{200, 100, 300},
		{-1, 0, 1},
		// make sure that the function works with negative numbers
		{-1, -1, 0},
		{-1, -20, 0},
		{-1, -20, -10},
	}

	for _, arg := range args {
		challenge.Function("Concat", student.BetweenUs, solutions.BetweenUs, arg[0], arg[1], arg[2])
	}
}
