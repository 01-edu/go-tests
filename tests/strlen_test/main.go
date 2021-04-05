package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{"Héllo!", random.Str(chars.Words, 26)}
	for i := 0; i < 10; i++ {
		args = append(args, random.Str(chars.Words, random.IntBetween(1, 20)))
	}
	for _, s := range args {
		challenge.Function("StrLen", student.StrLen, solutions.StrLen, s)
	}
}
