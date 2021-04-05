package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(random.IntSliceBetween(2, 20),
		random.IntBetween(-6, 20),
		-5,
		0,
	)
	for _, v := range args {
		challenge.Function("CollatzCountdown", student.CollatzCountdown, solutions.CollatzCountdown, v)
	}
}
