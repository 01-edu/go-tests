package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(rand.IntsBetween(-100, 150),
		20,
		0,
		9,
		2,
	)
	for _, arg := range table {
		challenge.Function("Fib", student.Fib, solutions.Fib, arg)
	}
}
