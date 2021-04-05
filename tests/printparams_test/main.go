package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := append(lib.MultRandWords(), "choumi is the best cat")
	for _, s := range table {
		challenge.Program("printparams", strings.Fields(s)...)
	}
}
