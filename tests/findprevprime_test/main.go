package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	elems := append(lib.MultRandIntBetween(0, 99999), 5, 4, 1)
	for _, elem := range elems {
		lib.Challenge("FindPrevPrime", student.FindPrevPrime, solutions.FindPrevPrime, elem)
	}
}
