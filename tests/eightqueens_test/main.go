package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	challenge.Function("EightQueens", student.EightQueens, solutions.EightQueens)
}
