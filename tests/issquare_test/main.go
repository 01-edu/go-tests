package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{1, 1},
		{2, 4},
		{3, 9},
		{4, 16},
		{5, 25},
		{6, 36},
		{7, 49},
		{-1, 0},
		{8, -64},
		{4, -16},
		{-1, 1},
		{-4, 16},
		{-5, 25},
	}
	for _, arg := range args {
		challenge.Function("IsSquare", student.IsSquare, solutions.IsSquare, arg[0], arg[1])
	}
}
