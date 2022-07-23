package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type arguments struct {
		ints []int
		init int
	}

	args := []arguments{
		{ints: []int{-1, 3}, init: 4},
		{ints: []int{-1, 3, 5}, init: 0},
		{ints: []int{-4, 3, 2}, init: -3},
		{ints: []int{}, init: -2},
		{ints: []int{-100, 300}, init: 4000},
		{ints: []int{-1, -1, 2, -2}, init: 5},
		{ints: []int{-1345, 3}, init: 10},
		{ints: []int{}, init: 3},
	}
	for _, arg := range args {
		challenge.Function("MultOrSum", student.MultOrSum, solutions.MultOrSum, arg.ints, arg.init)
	}
}
