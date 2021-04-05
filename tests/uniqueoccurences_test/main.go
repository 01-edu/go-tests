package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"abbaac",
		"ab",
		"abcacccazb",
		"",
	}
	for i := 0; i < 15; i++ {
		args = append(args, random.Str(chars.Lower, random.IntBetween(5, 10)))
	}
	for _, arg := range args {
		challenge.Program("uniqueoccurences", arg)
	}
}
