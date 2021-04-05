package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{rand.IntBetween(2, 20)}
	args = append(args, rand.IntsBetween(2, 20)...)
	args = append(args, rand.IntsBetween(2, 20)...)

	for _, v := range args {
		challenge.Function("ActiveBits", student.ActiveBits, solutions.ActiveBits, v)
	}
}
