package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		" ",
		"Hello 78 World!    4455 /",
	}
	for i := 0; i < 7; i++ {
		args = append(args, random.Str(chars.ASCII, random.IntBetween(5, 20)))
	}

	for _, arg := range args {
		challenge.Function("AlphaCount", student.AlphaCount, solutions.AlphaCount, arg)
	}
}
