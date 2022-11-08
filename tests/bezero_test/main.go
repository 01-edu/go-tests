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
		{-1, -1, 0},
		{-1, -20, 0},
		{-1, -20, -10},
		{1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
		{},
	}

	for _, arg := range args {
		challenge.Function("BeZero", student.BeZero, solutions.BeZero, arg)
	}
}
