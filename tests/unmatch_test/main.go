package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	i1 := random.IntBetween(-100, 100)
	i2 := random.IntBetween(-1000, 1000)
	i3 := random.IntBetween(-10, 10)
	args := [][]int{
		{1, 1, 2, 3, 4, 3, 4},
		{1, 1, 2, 4, 3, 4, 2, 3, 4},
		{1, 2, 1, 1, 4, 5, 5, 4, 1, 7},
		{1, 2, 3, 4, 5, 6, 7, 8},
		{0, 20, 91, 23, 10, 34},
		{1, 1, 2, 2, 3, 4, 3, 4, 5, 5, 8, 9, 8, 9},
		{i1, i2, i1, i2, i1 + i3, i1 + i3},
		{i1, i2, i1, i2, i1 + i3, i2 - i3},
	}
	for _, v := range args {
		challenge.Function("Unmatch", student.Unmatch, solutions.Unmatch, v)
	}
}
