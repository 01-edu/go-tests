package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]string{
		{"", "a", "b", "c"},
		{"a", "b", "c", "d"},
		{"Hello", "World", "!"},
		{"Hello", "World", "!"},
		{"Hello", ""},
		{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}
	for _, arg := range table {
		challenge.Function("AddFront", student.AddFront, solutions.AddFront, arg[0], arg[1:])
	}
}
