package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(random.StrSlice(chars.ASCII), "Hello! How are you?")
	for _, arg := range table {
		challenge.Function("ToUpper", student.ToUpper, solutions.ToUpper, arg)
	}
}
