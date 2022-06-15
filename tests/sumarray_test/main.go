package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{2, 5, 19},
		{2},
		{1, 2, 3, 5, 7, 24},
		{-10, 10, 20, 30, 40, 50},
		{},
		{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	for _, v := range args {
		challenge.Function("SumArray", student.SumArray, solutions.SumArray, v)
	}
}
