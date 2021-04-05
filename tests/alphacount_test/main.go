package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{
		" ",
		"Hello 78 World!    4455 /",
	}
	for l := 0; l < 7; l++ {
		a := lib.RandIntBetween(5, 20)
		b := lib.RandASCII()
		table = append(table, lib.RandStr(a, b))
	}

	for _, arg := range table {
		lib.Challenge("AlphaCount", student.AlphaCount, solutions.AlphaCount, arg)
	}
}
