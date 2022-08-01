package main

import (
	"student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int32{0, 5, 2, 3, 3, -9, 9, 4, 7, 10, 12, 1, 11, 6, 8, -6}

	for i := 0; i < len(args); i++ {
		challenge.Function("Y_quarter", student.Y_quarter, solutions.Y_quarter, args[i])
	}
}
