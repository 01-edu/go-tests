package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		rand.IntsBetween(0, rand.MaxInt),
		rand.MaxInt,
		321,
		321321,
		0,
	)
	for _, arg := range table {
		challenge.Function("PrintNbrInOrder", student.PrintNbrInOrder, solutions.PrintNbrInOrder, arg)
	}
}
