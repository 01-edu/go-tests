package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	args := []string{
		"you see it's easy to display the same thing",
		" only  it's   harder  ",
		"how   funny",
		"",
		rand.RandSpace(),
	}

	for _, v := range args {
		challenge.Program("cleanstr", v)
	}
	challenge.Program("cleanstr", "this is   not", "happening")
}
