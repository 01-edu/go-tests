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
		random.Str(allAscii, 2),
		random.Str(allAscii, 3),
		random.Str(allAscii, 4),
		random.Str(allAscii, 5),
		random.Str(allAscii, 6),
		random.Str(allAscii, 7),
		random.Str(allAscii, 8),
	}

	for _, arg := range args {
		challenge.Function("Ascii", student.Ascii, solutions.Ascii, arg)
	}
}
