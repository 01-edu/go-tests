package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(lib.MultRandIntBetween(2, 20),
		lib.RandIntBetween(-6, 20),
		-5,
		0,
	)
	for _, v := range args {
		lib.Challenge("CollatzCountdown", student.CollatzCountdown, solutions.CollatzCountdown, v)
	}
}
