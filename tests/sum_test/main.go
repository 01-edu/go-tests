package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]string{
		{"0", "0"},
		{"1", "0"},
		{"0", "1"},
		{"1", "2"},
		{"9", "9"},
		{"-9", "-9"},
		{"-3", "7"},
		{"8", "-7"},
		{"11", "7"},
		{"6", "256"},
	}
	for _, arg := range args {
		challenge.Function("Sum", student.Sum, solutions.Sum, arg[0], arg[1])
	}
}
