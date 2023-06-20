package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 13, -5}

	for _, arg := range args {
		challenge.Function("QuarterOfAYear", student.QuarterOfAYear, solutions.QuarterOfAYear, arg)
	}
}
