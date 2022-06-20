package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 3, 4, 5, 6, 7, 8, 9, 10},
		{-20, -10, 0, 10, 20},
		{29, 31, 32, 33, 34},
		{(int(^uint(0)>>1) * -1) - 1, int(^uint(0) >> 1)},
	}
	for _, arg := range args {
		challenge.Function("FindMissingNumber", student.FindMissingNumber, solutions.FindMissingNumber, arg)
	}
}
