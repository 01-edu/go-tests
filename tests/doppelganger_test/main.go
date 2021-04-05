package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

type node struct {
	big, little string
}

func main() {
	table := []node{}

	table = append(table,
		node{"aaaaaaa", "a"},
		node{"qwerty", "t"},
		node{"a", "b"},
	)

	for i := 0; i < 10; i++ {
		l := random.IntBetween(5, 30)
		big := random.RandStr(l, random.Lower)
		start := random.IntBetween(0, l-1)
		end := random.IntBetween(start+1, l-1)
		little := big[start:end]

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for i := 0; i < 10; i++ {
		big := random.RandStr(random.IntBetween(5, 30), random.Lower)
		little := random.RandStr(random.IntBetween(1, 29), random.Lower)

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for _, arg := range table {
		challenge.Function("DoppelGanger", student.DoppelGanger, solutions.DoppelGanger, arg.big, arg.little)
	}
}
