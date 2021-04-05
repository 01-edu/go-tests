package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{random.IntBetween(2, 20)}
	args = append(args, random.IntSliceBetween(2, 20)...)
	args = append(args, random.IntSliceBetween(2, 20)...)

	for _, v := range args {
		challenge.Function("ActiveBits", student.ActiveBits, solutions.ActiveBits, v)
	}
}
