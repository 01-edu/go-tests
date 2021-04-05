package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, arg := range table {
		challenge.Function("PrintCombN", student.PrintCombN, solutions.PrintCombN, arg)
	}
}
