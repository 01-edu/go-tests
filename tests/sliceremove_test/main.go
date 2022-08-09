package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][][]int{
		{{1, 2, 3, 4, 5}, {1}},
		{{1, 2, 3, 4, 5}, {5}},
		{{1, 2, 3, 4, 5}, {2}},
		{{}, {1}},
		{{1}, {1}},
		{{1, 3, 4}, {2}},
		{{1, 2}, {1}},
	}
	for _, s := range args {
		challenge.Function("SliceRemove", student.SliceRemove, solutions.SliceRemove, s[0], s[1][0])
	}
}
