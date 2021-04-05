package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := append(random.StrSlice(chars.Words),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ",
		"a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ",
	)
	for _, s := range table {
		challenge.Program("rot13", s)
	}
	challenge.Program("rot13", "1 argument", "2 arguments")
	challenge.Program("rot13", "1 argument", "2 arguments", "3 arguments")
}
