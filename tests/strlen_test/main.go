package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{"Héllo!", random.Words}
	for i := 0; i < 10; i++ {
		length := random.IntBetween(1, 20)
		args = append(args, random.RandStr(length, random.Words))
	}
	for _, s := range args {
		challenge.Function("StrLen", student.StrLen, solutions.StrLen, s)
	}
}

// TODO: refactor, simplify, no need for specific charset : check lib
