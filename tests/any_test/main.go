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

	args := []node{
		{
			f: is.Digit,
			a: []string{"Hello", "how", "are", "you"},
		},
		{
			f: is.Digit,
			a: []string{"This", "is", "4", "you"},
		},
	}

	for i := 0; i < 5; i++ {
		args = append(args, node{
			f: functions[random.IntBetween(0, len(functions)-1)],
			a: random.StrSlice(chars.Words),
		})
	}
	for i := 0; i < 5; i++ {
		args = append(args, node{
			f: is.Digit,
			a: random.StrSlice(chars.Digit),
		})
	}

	for i := 0; i < 5; i++ {
		args = append(args, node{
			f: is.Lower,
			a: random.StrSlice(chars.Lower),
		})
	}
	for i := 0; i < 5; i++ {
		args = append(args, node{
			f: is.Upper,
			a: random.StrSlice(chars.Upper),
		})
	}

	for _, arg := range args {
		challenge.Function("Any", student.Any, solutions.Any, arg.f, arg.a)
	}
}
