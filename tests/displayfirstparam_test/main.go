package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"1",
		"1 2",
		"1 2 3",
	)
	for _, s := range table {
		challenge.Program("displayfirstparam", strings.Fields(s)...)
	}
}
