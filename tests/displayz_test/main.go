package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := append(random.StrSlice(chars.Words),
		"dsfdz",
		"",
		"1",
		"1",
	)
	for _, s := range args {
		challenge.Program("displayz", strings.Fields(s)...)
	}

	challenge.Program("displayz", "1", "z")
}
