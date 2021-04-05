package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(random.IntSliceBetween(-100, 150),
		20,
		0,
		9,
		2,
	)
	for _, arg := range args {
		challenge.Function("Fib", student.Fib, solutions.Fib, arg)
	}
}
