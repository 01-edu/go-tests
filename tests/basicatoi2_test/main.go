package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		strconv.Itoa(random.MaxInt),
		"",
		"0",
		"Invalid123",
		"123Invalid",
		"Invalid",
		"1Invalid23",
		"123",
		"123.",
		"123.0",
	}
	for i := 0; i < 30; i++ {
		args = append(args, strconv.Itoa(random.IntBetween(0, random.MaxInt)))
	}

	for _, arg := range args {
		challenge.Function("BasicAtoi2", student.BasicAtoi2, solutions.BasicAtoi2, arg)
	}
}
