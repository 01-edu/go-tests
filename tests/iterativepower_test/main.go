package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	i := 0
	for i < 30 {
		nb := random.IntBetween(-8, 8)
		power := random.IntBetween(-10, 10)
		challenge.Function("IterativePower", student.IterativePower, solutions.IterativePower, nb, power)
		i++
	}
	challenge.Function("IterativePower", student.IterativePower, solutions.IterativePower, 0, 0)
	challenge.Function("IterativePower", student.IterativePower, solutions.IterativePower, 0, 1)
}
