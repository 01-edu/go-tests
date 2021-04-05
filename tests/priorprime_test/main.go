package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(random.IntSliceBetween(0, 1000),
		50,
		13,
		10,
		0,
		1,
		2,
	)
	for _, arg := range args {
		challenge.Function("PriorPrime", student.PriorPrime, solutions.PriorPrime, arg)
	}
}
