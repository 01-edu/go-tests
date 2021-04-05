package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := append(lib.MultRandWords(),
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
