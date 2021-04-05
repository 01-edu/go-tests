package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		"",
		"abcdefghijklm",
		"the time of contempt precedes that of indifference",
		"he stared at the mountain",
		"qw qw e qwsa d",
	}

	// 3 valid random sentences with no spaces at the beginning nor the end and only one space for separator.
	for i := 0; i < 3; i++ {
		numberOfWords := random.IntBetween(1, 6)
		sentence := random.RandAlnum()
		for j := 0; j < numberOfWords; j++ {
			sentence += " " + random.RandAlnum()
		}
		sentence += random.RandAlnum()
		table = append(table, sentence)
	}

	for _, s := range table {
		challenge.Program("revwstr", s)
	}

	challenge.Program("revwstr")
	challenge.Program("revwstr", "1param", "2param", "3param", "4param")
}
