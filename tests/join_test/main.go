package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	seps := []string{" ", "-", " ,", "_", "SPC", " . "}

	args := [][]string{
		{"hello", "how", "are", "you", "doing"},
		{"fine", "and", "you"},
		{"I'm", "O.K."},
	}

	for i := 0; i < 5; i++ {
		// random position for the slice of arguments
		posA := random.IntBetween(0, len(args)-1)
		// random position for the slice of separators
		posS := random.IntBetween(0, len(seps)-1)

		challenge.Function("Join", student.Join, solutions.Join, args[posA], seps[posS])
	}
}
