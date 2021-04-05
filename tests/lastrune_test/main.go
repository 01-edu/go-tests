package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := rand.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		rand.RandStr(rand.IntBetween(1, 15), rand.RandAlnum()),
	)
	for _, arg := range table {
		challenge.Function("LastRune", student.LastRune, solutions.LastRune, arg)
	}
}
