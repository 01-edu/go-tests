package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	arr := []string{
		"6H8a",
		"3p6i1W",
		"2O5u2H0e",
		"2d",
		"2a 6p8f",
		"2p4;7g",
		"2t4dD",
		"82t4D",
		"2k5%9.",
		"",
		"5w6	2t",
		"8o2d7-1a",
		"4	1w8*6x",
		"2p6t5z",
		"7w1mw0q",
		"0o5d3j",
		"3a2b1c",
	}
	for _, s := range arr {
		challenge.Function("Unzipstring", student.Unzipstring, solutions.Unzipstring, s)
	}
}
