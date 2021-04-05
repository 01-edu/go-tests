package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][2]int{
		{50, 43},
		{13, 13},
		{10, 9},
		{5, 9},
	}
	for i := 0; i < 15; i++ {
		table = append(table, [2]int{
			random.IntBetween(0, 30),
			random.IntBetween(0, 100),
		})
	}
	for _, arg := range table {
		challenge.Function("Volumechanger", student.Volumechanger, solutions.Volumechanger, arg[0], arg[1])
	}
}
