package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []int{1, 2, 3, 4, 9285, 74584, 0, -5585, 75418, 99999, -1, 526, 36}
	for _, n := range table {
		challenge.Function("Divisors", student.Divisors, solutions.Divisors, n)
	}
}
