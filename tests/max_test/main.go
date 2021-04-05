package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{lib.RandInt()}
	limit := lib.RandIntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, lib.RandInt())
	}

	lib.Challenge("Max", student.Max, solutions.Max, args)
}
