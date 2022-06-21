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
	}
	for _, s := range arr {
		challenge.Function("Unzipstring", student.Unzipstring, solutions.Unzipstring, s)
	}
}
