package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello World!")
	for _, arg := range table {
		challenge.Function("PrintStr", student.PrintStr, solutions.PrintStr, arg)
	}
}
