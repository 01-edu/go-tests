package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]int{
		{1, 1},
		{0, 0},
		{4, 2},
		{7, 3},
		{2, -1},
		{-1, 1},
		{3, 4},
		{-1, -1},
		{100, 303},
		{2147483647, -2147483648},
	}
	for _, arg := range table {
		challenge.Function("BinaryAddition", student.BinaryAddition, solutions.BinaryAddition, arg)
	}
}
