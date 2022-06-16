package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {

	args:= [][]int{
		{1, 1},
		{2, 4},
		{3, 9},
		{4, 16},
		{5, 25},
		{6, 36},
		{7, 49},
		{-1, 0},
	}
	for _, arg := range args {
		challenge.Function("IsTheSquareAChild", student.IsTheSquareAChild, solutions.IsTheSquareAChild, arg[0], arg[1])
	}
}
