package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(random.IntSliceBetween(0, 25),
		4,
		5,
		-5,
	)
	for _, arg := range args {
		challenge.Function("Fibonacci", student.Fibonacci, solutions.Fibonacci, arg)
	}
}
