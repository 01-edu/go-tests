package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.MultRandWords(),
		"dsfdz",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		challenge.Program("displayz", strings.Fields(s)...)
	}

	challenge.Program("displayz", "1", "z")
}
