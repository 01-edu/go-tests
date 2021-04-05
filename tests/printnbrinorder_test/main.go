package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		random.IntSliceBetween(0, random.MaxInt),
		random.MaxInt,
		321,
		321321,
		0,
	)
	for _, arg := range table {
		challenge.Function("PrintNbrInOrder", student.PrintNbrInOrder, solutions.PrintNbrInOrder, arg)
	}
}
