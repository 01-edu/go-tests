package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{0, 5, 2, 3, 3, -9, 9, 4, 7, 10, 12, 1, 11, 6, 8, -6}

	for _, arg := range args {
		challenge.Function("QuarterOfAYear", student.QuarterOfAYear, solutions.QuarterOfAYear, arg)
	}
}
