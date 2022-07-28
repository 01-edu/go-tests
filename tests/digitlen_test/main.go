package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]int{
		{10, 2},
		{10, 3},
		{10, 4},
		{10, 5},
		{10, 6},
		{100, -1},
		{100, 10},
		{-100, 10},
		{-100, -1},
		{-100, 2},
		{-100, 15},
		{-100, 8},
	}

	for _, arg := range table {
		challenge.Function("DigitLen", student.DigitLen, solutions.DigitLen, arg[0], arg[1])
	}
}
