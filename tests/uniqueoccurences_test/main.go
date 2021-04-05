package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := []string{
		"abbaac",
		"ab",
		"abcacccazb",
		"",
	}

	for i := 0; i < 15; i++ {
		table = append(table, lib.RandStr(lib.RandIntBetween(5, 10), lib.Lower))
	}

	for _, arg := range table {
		lib.Program("uniqueoccurences", arg)
	}
}
