package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.MultRandWords(), "choumi is the best cat")
	for _, s := range table {
		challenge.Program("printparams", strings.Fields(s)...)
	}
}
