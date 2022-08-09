package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][][]int{
		{{1, 2, 3}, {1}},
		{{3}, {5}},
		{{1, 2}, {6}},
		{{}, {0}},
		{{12, 13, 32, 23, 3, 3, 2, 2, 1, 33, 4}, {-10}},
		{{}, {-10}},
	}
	for _, arg := range table {
		challenge.Function("SliceAdd", student.SliceAdd, solutions.SliceAdd, arg[0], arg[1][0])
	}
}
