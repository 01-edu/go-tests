package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{}, {-10, 0, 2}},
		{{-10, 0, 2}, {}},
		{{}, {}},
		{{1, 2, 3}, {4, 5, 6, 3, 4, 5, 6}},
		{{0, 0, 0}, {0, 0, 0}},
	}

	for _, arg := range args {
		challenge.Function("ConcatSlice", student.ConcatSlice, solutions.ConcatSlice, arg[0], arg[1])
	}
}
