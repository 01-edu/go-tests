package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		plus  int
		minus int
		rand  int
	}
	table := []node{
		{50, 43, 20},
		{13, 13, 0},
		{10, 9, 0},
		{5, 9, 2},
	}

	for i := 0; i < 15; i++ {
		table = append(table, node{
			plus:  random.IntBetween(0, 10),
			minus: random.IntBetween(0, 10),
			rand:  random.IntBetween(0, 10),
		})
	}
	for _, arg := range table {
		challenge.Function("Nauuo", student.Nauuo, solutions.Nauuo, arg.plus, arg.minus, arg.rand)
	}
}
