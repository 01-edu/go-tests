package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
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
		l := rand.IntBetween(5, 30)
		big := rand.RandStr(l, rand.Lower)
		start := rand.IntBetween(0, l-1)
		end := rand.IntBetween(start+1, l-1)
		little := big[start:end]

		value := node{
			big:    big,
			little: little,
		}

		table = append(table, value)
	}

	for i := 0; i < 10; i++ {
		big := rand.RandStr(rand.IntBetween(5, 30), rand.Lower)
		little := rand.RandStr(rand.IntBetween(1, 29), rand.Lower)

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
