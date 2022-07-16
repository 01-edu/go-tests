package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]int{
		{10, 10, 3, 1, 2, 3, 5, 3, 63},
		{-10, -10, 03, 1, 2, 3, 5, 3, 63},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 3, -100, 39, 4, 10, -29, -49, 92},
		{-1},
		{23},
		{3, -1, -2, -3},
		{0, -3, -2, -1},
		{},
	}
	for _, arg := range table {
		challenge.Function("CountNegative", student.CountNegative, solutions.CountNegative, arg)
	}
}
