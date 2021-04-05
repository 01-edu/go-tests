package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(
		random.StrSlice(chars.ASCII),
		"Hello! How are you? How+are+things+4you?",
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"9a",
		"9a LALALA!",
	)
	for _, arg := range args {
		challenge.Function("Capitalize", student.Capitalize, solutions.Capitalize, arg)
	}
}
