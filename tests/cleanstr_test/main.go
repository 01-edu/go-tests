package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := []string{
		"you see it's easy to display the same thing",
		" only  it's   harder  ",
		"how   funny",
		"",
		strings.Repeat(" ", 13),
	}

	for _, v := range args {
		challenge.Program("cleanstr", v)
	}
	challenge.Program("cleanstr", "this is   not", "happening")
}
