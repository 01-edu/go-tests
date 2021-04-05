package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []int{
		20,
		1,
		9,
		2,
	}
	table = append(table, random.IntSliceBetween(1, 1500)...)

	for _, arg := range table {
		challenge.Function("InterestingNumber", student.InterestingNumber, solutions.InterestingNumber, arg)
	}
}
