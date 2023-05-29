package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := make([]string, 30)
	for i := range args {
		args[i] = strconv.Itoa(random.IntBetween(0, random.MaxInt))
	}
	args = append(args,
		strconv.Itoa(random.MaxInt),
		"0",
		"12345",
		"0000012345",
		"000000",
	)
	for _, arg := range args {
		challenge.Function("BasicAtoi", student.BasicAtoi, solutions.BasicAtoi, arg)
	}
}
