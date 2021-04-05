package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"UD",
		"LL",
	}

	for i := 0; i < 15; i++ {
		args = append(args, random.Str("UDLR", random.IntBetween(5, 1000)))
	}

	for _, arg := range args {
		challenge.Program("robottoorigin", arg)
	}
}
