package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]string{
		{},
		{"A", "b", "C", "d"},
		{"HA", "HI", "UH"},
		{"A", "BE", "c", "DE"},
		{"two", "words", "or", "three"},
		{"café", "fête", "naïve", "façade"},
		{
			"Lorem ipsum dolor sit amet,",
			" consectetur adipiscing elit,",
			"sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		},
	}

	for _, arg := range args {
		challenge.Function("EvenLength", student.EvenLength, solutions.EvenLength, arg)
	}
}
