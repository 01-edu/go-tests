package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		"",
		"             a as",
		"   f     d",
		"   asd    ad",
		"   salut !!! ",
		" salut ! ! !",
		"salut ! !",
	)
	for _, s := range table {
		challenge.Program("firstword", s)
	}
}
