package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]int{
		{10, 10},
		{-10, -10},
		{0, 0},
		{1, 1},
		{-1, 20},
		{23, 2},
		{3, 0},
		{0, 39},
	}
	for _, arg := range table {
		challenge.Function("AddIfPositive", student.AddIfPositive, solutions.AddIfPositive, arg[0], arg[1])
	}
}
