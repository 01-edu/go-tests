package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := append(rand.MultRandWords(),
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
