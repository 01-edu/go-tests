package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	arg := rand.Ints()
	arg = append(arg, rand.Int())
	for i := 0; i < 15; i++ {
		challenge.Function("Abort", student.Abort, solutions.Abort, arg[0], arg[1], arg[2], arg[3], arg[4])
		arg = rand.Ints()
		arg = append(arg, rand.Int())
	}
}
