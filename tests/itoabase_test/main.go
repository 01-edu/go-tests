package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 30; i++ {
		value := lib.RandIntBetween(-1000000, 1000000)
		base := lib.RandIntBetween(2, 16)
		lib.Challenge("ItoaBase", student.ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := lib.RandIntBetween(2, 16)
		lib.Challenge("ItoaBase", student.ItoaBase, solutions.ItoaBase, lib.MaxInt, base)
		lib.Challenge("ItoaBase", student.ItoaBase, solutions.ItoaBase, lib.MinInt, base)
	}
}
