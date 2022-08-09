package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 5, 4, 4, 2, 1},
		{1, 1, 1, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
		{},
		{1},
		{0, 0, 0, 2},
		{10, 20, 10, 20, 30, 20, 10, 20, 10, 20},
	}
	for _, v := range table {
		challenge.Function("RemoveDuplicate", solutions.RemoveDuplicate, student.RemoveDuplicate, v)
	}
}
