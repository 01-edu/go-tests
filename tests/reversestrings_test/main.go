package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]string{
		{"good", "morning!"},
		{"Zoo", "mice"},
		{"", ""},
		{"Even", "more"},
		{"Try", "this", "one"},
		{"A", "B", "C", "D"},
		{"1", "3", "$", "301", "hELLO", "WORLD"},
	}
	for _, arg := range args {
		challenge.Function("ReverseStrings", student.ReverseStrings, solutions.ReverseStrings, arg)
	}
}
