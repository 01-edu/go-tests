package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []int{
		5, 3,
		5, 1,
		1, 1,
		1, 5,
	}

	table = append(table,
		0, 0,
		-1, 6,
		6, -1,
		lib.RandIntBetween(1, 20), lib.RandIntBetween(1, 20),
	)

	for i := 0; i < len(table); i += 2 {
		if i != len(table)-1 {
			lib.Challenge("QuadC", student.QuadC, solutions.QuadC, table[i], table[i+1])
		}
	}
}
