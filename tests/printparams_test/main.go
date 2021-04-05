package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words), "choumi is the best cat")
	for _, s := range table {
		challenge.Program("printparams", strings.Fields(s)...)
	}
}
