package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
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
			lib.RandIntBetween(0, 30),
			lib.RandIntBetween(0, 100),
		})
	}
	for _, arg := range table {
		lib.Function("Volumechanger", student.Volumechanger, solutions.Volumechanger, arg[0], arg[1])
	}
}
