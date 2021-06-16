package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func randSliceInt() [10]int {
	var ints [10]int
	for i := range ints {
		ints[i] = random.IntBetween(0, 1000)
	}
	return ints
}

func main() {
	table := [][10]int{{104, 101, 108, 108, 111, 16, 21, 42}}

	for i := 0; i < 5; i++ {
		table = append(table, randSliceInt())
	}

	for _, t := range table {
		challenge.Function("PrintMemory", student.PrintMemory, solutions.PrintMemory, t)
	}
}
