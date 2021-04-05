package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(rand.IntBetween(0, rand.MaxInt))
	}
	table = append(table,
		strconv.Itoa(rand.MaxInt),
		"",
		"0",
		"12345",
		"0000012345",
		"000000",
	)
	for _, arg := range table {
		challenge.Function("BasicAtoi", student.BasicAtoi, solutions.BasicAtoi, arg)
	}
}
