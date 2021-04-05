package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 30; i++ {
		value := rand.IntBetween(-1000000, 1000000)
		base := rand.IntBetween(2, 16)
		challenge.Function("ItoaBase", student.ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := rand.IntBetween(2, 16)
		challenge.Function("ItoaBase", student.ItoaBase, solutions.ItoaBase, rand.MaxInt, base)
		challenge.Function("ItoaBase", student.ItoaBase, solutions.ItoaBase, rand.MinInt, base)
	}
}
