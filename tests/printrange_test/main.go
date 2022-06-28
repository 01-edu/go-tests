package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

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
		{-10, 15},
		{10, -15},
	}
	for _, arg := range args {
		challenge.Function("PrintRange", student.PrintRange, solutions.PrintRange, arg[0], arg[1])
	}
}
