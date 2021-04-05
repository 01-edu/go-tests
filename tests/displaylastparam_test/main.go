package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		" ",
		"1",
		"1 2",
		"1 2 3",
	)
	for _, s := range table {
		challenge.Program("displaylastparam", strings.Fields(s)...)
	}
}
