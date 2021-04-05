package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/is"
	"github.com/01-edu/go-tests/solutions"
)

func isPositive(i int) bool {
	return i > 0
}

func isNegative(i int) bool {
	return i < 0
}

func main() {
	functions := []func(int) bool{isPositive, isNegative, is.Prime}

	type node struct {
		f func(int) bool
		a []int
	}

	table := []node{}

	for i := 0; i < 15; i++ {
		function := functions[lib.RandIntBetween(0, len(functions)-1)]
		val := node{
			f: function,
			a: lib.MultRandIntBetween(-1000000, 1000000),
		}
		table = append(table, val)
	}

	table = append(table, node{
		f: is.Prime,
		a: []int{1, 2, 3, 4, 5, 6},
	})

	for _, arg := range table {
		lib.Function("Map", student.Map, solutions.Map, arg.f, arg.a)
	}
}
