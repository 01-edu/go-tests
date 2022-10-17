package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int32{0, 6, 33, -3, 5, 8, -84, 938, 38, 92, -78, -53, -23, 33, 4, 64, 76, 24, 90, 22, 31, 241, 51, 391}
	for i := 0; i < len(args); i++ {
		challenge.Function("BinaryCheck", student.BinaryCheck, solutions.BinaryCheck, args[i])
	}
}
