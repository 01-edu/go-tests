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
		s, substr string
	}

	args := []node{
		{"aaaaaaa", "a"},
		{"qwerty", "t"},
		{"a", "b"},
	}

	for i := 0; i < 10; i++ {
		l := random.IntBetween(5, 30)
		s := random.Str(chars.Lower, l)
		start := random.IntBetween(0, l-1)
		end := random.IntBetween(start+1, l-1)
		substr := s[start:end]
		args = append(args, node{s, substr})
	}

	for i := 0; i < 10; i++ {
		args = append(args, node{
			s:      random.Str(chars.Lower, random.IntBetween(5, 30)),
			substr: random.Str(chars.Lower, random.IntBetween(1, 29)),
		})
	}

	for _, arg := range args {
		challenge.Function("DoppelGanger", student.DoppelGanger, solutions.DoppelGanger, arg.s, arg.substr)
	}
}
