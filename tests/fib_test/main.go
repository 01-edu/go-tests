package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(lib.MultRandIntBetween(-100, 150),
		20,
		0,
		9,
		2,
	)
	for _, arg := range table {
		lib.Function("Fib", student.Fib, solutions.Fib, arg)
	}
}
