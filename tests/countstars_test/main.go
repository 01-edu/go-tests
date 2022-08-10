package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []int{1, 2, 9, -258, 0, 100, 48, 94, 41, -56, 520}
	for _, num := range table {
		challenge.Function("CountStars", student.CountStars, solutions.CountStars, num)
	}
}
