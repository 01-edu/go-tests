package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	intsArgs := [][]int{
		{},
		{},
		{1, 2, 3, 4, 5},
		{1, 2},
		{23, 25, 26, 65},
		{98, 345, 456},
	}
	positionArgs := []int{0, 1, 6, 1, 3, 2}
	for index, arg := range intsArgs {
		challenge.Function("Delete", student.Delete, solutions.Delete, arg, positionArgs[index])
	}
}
