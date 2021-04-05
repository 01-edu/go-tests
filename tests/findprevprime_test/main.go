package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	elems := append(rand.IntsBetween(0, 99999), 5, 4, 1)
	for _, elem := range elems {
		challenge.Function("FindPrevPrime", student.FindPrevPrime, solutions.FindPrevPrime, elem)
	}
}
