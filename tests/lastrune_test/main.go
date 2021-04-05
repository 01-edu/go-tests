package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := lib.MultRandASCII()
	table = append(table,
		"Hello!",
		"Salut!",
		"Ola!",
		lib.RandStr(lib.RandIntBetween(1, 15), lib.RandAlnum()),
	)
	for _, arg := range table {
		lib.Function("LastRune", student.LastRune, solutions.LastRune, arg)
	}
}
