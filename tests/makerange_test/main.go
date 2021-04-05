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

	table := []node{}

	// 15 random pairs of ints for a Valid Range
	for i := 0; i < 15; i++ {
		minVal := random.IntBetween(-10000000, 1000000)
		gap := random.IntBetween(1, 20)
		val := node{
			min: minVal,
			max: minVal + gap,
		}
		table = append(table, val)
	}
	// 15 random pairs of ints with ||invalid range||
	for i := 0; i < 15; i++ {
		minVal := random.IntBetween(-10000000, 1000000)
		gap := random.IntBetween(1, 20)
		val := node{
			min: minVal,
			max: minVal - gap,
		}
		table = append(table, val)
	}

	table = append(table,
		node{min: 0, max: 1},
		node{min: 0, max: 0},
		node{min: 5, max: 10},
		node{min: 10, max: 5},
	)
	for _, arg := range table {
		challenge.Function("MakeRange", student.MakeRange, solutions.MakeRange, arg.min, arg.max)
	}
}
