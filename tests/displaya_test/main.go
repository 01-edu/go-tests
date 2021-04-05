package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := append(lib.MultRandWords(),
		"dsfda",
		"",
		"1",
		"1",
	)
	for _, s := range table {
		lib.Program("displaya", strings.Fields(s)...)
	}
	lib.Program("displaya", "1", "a")
}
