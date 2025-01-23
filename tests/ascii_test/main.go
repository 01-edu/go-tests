package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	allAscii := ""
	for i := 0; i <= 255; i++ {
		allAscii += string(i)
	}
	args := []string{
		"",
		random.Str(allAscii, 1),
		random.Str(allAscii, 8),
		"Instead of this many random tests leave just a few and add some tests with long strings and multiple spaces",
		"one two  three   four    ",
	}

	for _, arg := range args {
		challenge.Function("Ascii", student.Ascii, solutions.Ascii, arg)
	}
}
