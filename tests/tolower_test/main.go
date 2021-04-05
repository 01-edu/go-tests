package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(random.MultRandASCII(), "Hello! How are you?")
	for _, arg := range table {
		challenge.Function("ToLower", student.ToLower, solutions.ToLower, arg)
	}
}
