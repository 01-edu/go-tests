package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := random.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		random.RandStr(random.IntBetween(1, 15), random.RandAlnum()),
	)
	for _, arg := range table {
		challenge.Function("LastRune", student.LastRune, solutions.LastRune, arg)
	}
}
