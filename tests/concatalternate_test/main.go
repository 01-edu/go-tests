package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{1, 2, 3}, {4, 5}},
		{{}, {4, 5, 6}},
		{{1, 2, 3}, {}},
		{{}, {}},
		{{1, 2, 4}, {10, 20, 30, 40, 50}},
	}
	for _, arg := range args {
		challenge.Function("ConcatAlternate", student.ConcatAlternate, solutions.ConcatAlternate, arg[0], arg[1])
	}
}
