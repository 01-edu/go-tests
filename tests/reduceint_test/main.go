package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	f := []func(int, int) int{
		func(accumulator, currentValue int) int {
			return accumulator + currentValue
		},
		func(accumulator, currentValue int) int {
			return accumulator - currentValue
		},
		func(accumulator, currentValue int) int {
			return currentValue * accumulator
		},
	}
	var argInt []int

	type node struct {
		a         []int
		functions []func(int, int) int
	}

	table := []node{}
	for i := 0; i < 4; i++ {
		argInt = random.IntSliceBetween(0, 50)
		table = append(table, node{
			a:         argInt,
			functions: f,
		})
	}

	for _, v := range table {
		for _, f := range v.functions {
			challenge.Function("ReduceInt", student.ReduceInt, solutions.ReduceInt, v.a, f)
		}
	}
}
