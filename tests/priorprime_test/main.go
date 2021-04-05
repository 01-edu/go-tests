package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(lib.MultRandIntBetween(0, 1000),
		50,
		13,
		10,
		0,
		1,
		2,
	)
	for _, arg := range table {
		lib.Function("PriorPrime", student.PriorPrime, solutions.PriorPrime, arg)
	}
}
