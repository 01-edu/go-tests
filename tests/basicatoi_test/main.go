package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(random.IntBetween(0, random.MaxInt))
	}
	table = append(table,
		strconv.Itoa(random.MaxInt),
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
