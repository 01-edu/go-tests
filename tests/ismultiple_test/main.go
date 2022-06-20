package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {

	table := []int{
		10,
		9,
		8,
		0,
		-10,
		-9,
		2948483,
		4848,
		3282,
		-3282,
	}
	for _, num := range table {
		challenge.Function("IsMultiple", student.IsMultiple, solutions.IsMultiple, num)
	}
}
