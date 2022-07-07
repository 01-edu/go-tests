package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	arr := []int{1, 244, 3, 45, -58, 8, 2147483647, 0, 98, 7}
	for _, arg := range arr {
		challenge.Function("SquareRoot", student.SquareRoot, solutions.SquareRoot, arg)
	}
}
