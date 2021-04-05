package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"dsfdz",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		lib.Program("displayz", strings.Fields(s)...)
	}

	lib.Program("displayz", "1", "z")
}
