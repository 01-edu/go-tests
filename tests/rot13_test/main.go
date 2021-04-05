package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ",
		"a b c d e f g h ijklmnopqrstuvwxyz A B C D E FGHIJKLMNOPRSTUVWXYZ",
	)
	for _, s := range table {
		lib.Program("rot13", s)
	}
	lib.Program("rot13", "1 argument", "2 arguments")
	lib.Program("rot13", "1 argument", "2 arguments", "3 arguments")
}
