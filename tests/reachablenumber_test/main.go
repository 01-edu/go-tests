package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{
		20,
		1,
		9,
		2,
	}
	for i := 0; i < 25; i++ {
		args = append(args, random.IntSliceBetween(1, 877)...)
	}
	for _, arg := range args {
		challenge.Function("ReachableNumber", student.ReachableNumber, solutions.ReachableNumber, arg)
	}
}
