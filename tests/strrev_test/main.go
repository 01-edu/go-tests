package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		lib.MultRandASCII(),
		"Hello!",
		"Bonjour!",
		"Hola!",
	)
	for _, arg := range table {
		challenge.Function("StrRev", student.StrRev, solutions.StrRev, arg)
	}
}
