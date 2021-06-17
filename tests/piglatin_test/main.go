package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"", "pig", "is", "crunch", "crnch", "something else",
	}

	for i := 0; i < 4; i++ {
		args = append(args, random.Str(chars.Basic, 7))
	}

	for _, v := range args {
		challenge.Program("piglatin", v)
	}
}
