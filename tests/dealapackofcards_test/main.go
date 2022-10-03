package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		{11, 2, 14, 4, 5, 17, 7, 8, 9, 10, 11, 12},
		{1, 2, 3, 1, 5, 5, 7, 8, 9, 19, 16, 12}}

	for _, s := range table {
		challenge.Function("DealAPackOfCards", student.DealAPackOfCards, solutions.DealAPackOfCards, s)
	}
}
