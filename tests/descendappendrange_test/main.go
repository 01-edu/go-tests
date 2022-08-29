package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		min int
		max int
	}
	args := []node{
		{min: 0, max: 1},
		{min: 0, max: 0},
		{min: 5, max: 10},
		{min: 10, max: 5},
	}

	// 15 random pairs of ints for a Valid Range
	for i := 0; i < 15; i++ {
		minVal := random.IntBetween(-10000000, 1000000)
		gap := random.IntBetween(1, 20)
		args = append(args, node{
			min: minVal,
			max: minVal + gap,
		})
	}

	// 15 random pairs of ints with ||invalid range||
	for i := 0; i < 15; i++ {
		minVal := random.IntBetween(-10000000, 1000000)
		gap := random.IntBetween(1, 20)
		args = append(args, node{
			min: minVal,
			max: minVal - gap,
		})
	}

	for _, arg := range args {
		challenge.Function("DescendAppendRange", student.DescendAppendRange, solutions.DescendAppendRange, arg.min, arg.max)
	}
}
