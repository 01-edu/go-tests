package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{rand.Int()}
	limit := rand.IntBetween(20, 50)
	for i := 0; i < limit; i++ {
		args = append(args, rand.Int())
	}

	challenge.Function("Max", student.Max, solutions.Max, args)
}
