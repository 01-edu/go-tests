package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	table := []string{
		"abbaac",
		"ab",
		"abcacccazb",
		"",
	}

	for i := 0; i < 15; i++ {
		table = append(table, rand.RandStr(rand.IntBetween(5, 10), rand.Lower))
	}

	for _, arg := range table {
		challenge.Program("uniqueoccurences", arg)
	}
}
