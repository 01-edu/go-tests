package main

import (
	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcdefghi jklmnop qrstuvwxyz ABCDEFGHI JKLMNOPQR STUVWXYZ ! ",
	)
	for _, s := range table {
		challenge.Program("switchcase", s)
	}
}
