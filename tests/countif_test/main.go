package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/is"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	functions := []func(string) bool{is.Digit, is.Lower, is.Upper}

	type node struct {
		f func(string) bool
		a []string
	}

	table := []node{}

	for i := 0; i < 5; i++ {
		function := functions[random.IntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: random.StrSlice(chars.Words),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Digit,
			a: random.StrSlice(chars.Digit),
		})
	}

	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Lower,
			a: random.StrSlice(chars.Lower),
		})
	}
	for i := 0; i < 5; i++ {
		table = append(table, node{
			f: is.Upper,
			a: random.StrSlice(chars.Upper),
		})
	}

	table = append(table,
		node{
			f: is.Digit,
			a: []string{"Hello", "how", "are", "you"},
		},
		node{
			f: is.Digit,
			a: []string{"This", "is", "4", "you"},
		},
	)

	for _, arg := range table {
		challenge.Function("CountIf", student.CountIf, solutions.CountIf, arg.f, arg.a)
	}
}
