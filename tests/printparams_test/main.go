package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := append(rand.MultRandWords(), "choumi is the best cat")
	for _, s := range table {
		challenge.Program("printparams", strings.Fields(s)...)
	}
}
