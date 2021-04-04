package main

import (
	"strings"

	student "student"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	table := append(lib.MultRandASCII(), "Hello! How are you?")
	for _, arg := range table {
		lib.Challenge("ToUpper", student.ToUpper, strings.ToUpper, arg)
	}
}
