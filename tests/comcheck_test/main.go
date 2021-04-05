package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		"01",
		"galaxy",
		"galaxy01",
		" 01",
		"as das d 01 asd",
		"as galaxy d 12",
		"as ds galaxy 01 asd")

	for _, s := range table {
		challenge.Program("comcheck", strings.Fields(s)...)
	}
}
