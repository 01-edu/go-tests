package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		lib.MultRandIntBetween(0, lib.MaxInt),
		lib.MaxInt,
		321,
		321321,
		0,
	)
	for _, arg := range table {
		challenge.Function("PrintNbrInOrder", student.PrintNbrInOrder, solutions.PrintNbrInOrder, arg)
	}
}
