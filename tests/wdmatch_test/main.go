package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := append(lib.MultRandWords(),
		" ",
		"faya fgvvfdxcacpolhyghbreda",
		"faya fgvvfdxcacpolhyghbred",
		"error rrerrrfiiljdfxjyuifrrvcoojh",
		"error rrerrrfiiljdfxjyuifrrvcoojrh",
	)

	for _, s := range table {
		challenge.Program("wdmatch", strings.Fields(s)...)
	}
}
