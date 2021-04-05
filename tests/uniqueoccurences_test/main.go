package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		"abbaac",
		"ab",
		"abcacccazb",
		"",
	}

	for i := 0; i < 15; i++ {
		table = append(table, random.RandStr(random.IntBetween(5, 10), random.Lower))
	}

	for _, arg := range table {
		challenge.Program("uniqueoccurences", arg)
	}
}
