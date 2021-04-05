package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		s         string
		toCompare string
	}

	args := []node{
		{s: "Hello!", toCompare: "Hello!"},
		{s: "Salut!", toCompare: "lut!"},
		{s: "Ola!", toCompare: "Ol"},
	}

	// the first 15 values are returning 0 for this test
	for i := 0; i < 15; i++ {
		wordToTest := random.Str(chars.ASCII, 13)
		args = append(args, node{
			s:         wordToTest,
			toCompare: wordToTest,
		})
	}

	// the next 15 values are supposed to return 1 or -1 for this test
	for i := 0; i < 15; i++ {
		args = append(args, node{
			s:         random.Str(chars.ASCII, 13),
			toCompare: random.Str(chars.ASCII, 13),
		})
	}

	for _, arg := range args {
		challenge.Function("Compare", student.Compare, solutions.Compare, arg.s, arg.toCompare)
	}
}
