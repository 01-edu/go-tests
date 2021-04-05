package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := append(rand.MultRandWords(), "1 2 3 4 5")

	for _, s := range table {
		challenge.Program("boolean", strings.Fields(s)...)
	}
}
