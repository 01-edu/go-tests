package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int32{20, 15, -125, 10, 5, 4, -9, 9, 24, 0, -85, 6, 11}

	for i := 0; i < len(args); i++ {
		challenge.Function("BuzZinga", student.BuzZinga, solutions.BuzZinga, args[i])
	}
}
