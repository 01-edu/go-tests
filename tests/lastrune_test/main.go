package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := append(random.StrSlice(chars.ASCII),
		"Hello!",
		"Salut!",
		"Ola!",
		random.Str(chars.Alnum, random.IntBetween(1, 15)),
	)
	for _, arg := range args {
		challenge.Function("LastRune", student.LastRune, solutions.LastRune, arg)
	}
}
