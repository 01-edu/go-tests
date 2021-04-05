package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []int{
		20,
		1,
		9,
		2,
	}
	table = append(table, lib.MultRandIntBetween(1, 1500)...)

	for _, arg := range table {
		lib.Function("InterestingNumber", student.InterestingNumber, solutions.InterestingNumber, arg)
	}
}
