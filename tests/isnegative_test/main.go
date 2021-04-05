package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		random.IntSlice(),
		random.MinInt,
		random.MaxInt,
		0,
	)
	for _, arg := range table {
		challenge.Function("IsNegative", student.IsNegative, solutions.IsNegative, arg)
	}
}
