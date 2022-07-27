package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{1, 2},
		{1, 10},
		{1, 100},
		{100, 1},
		{99, 1},
		{1, 1},
		{1, 0},
		{0, 1},
		{-1, 10},
		{1, 99},
		{99, -1},
	}
	for _, arg := range args {
		challenge.Function("FromTo", student.FromTo, solutions.FromTo, arg[0], arg[1])
	}
}
