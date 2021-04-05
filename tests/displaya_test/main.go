package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		"dsfda",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		challenge.Program("displaya", strings.Fields(s)...)
	}
	challenge.Program("displaya", "1", "a")
}
