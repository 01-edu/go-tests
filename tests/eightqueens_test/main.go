package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	lib.Challenge("EightQueens", student.EightQueens, solutions.EightQueens)
}
