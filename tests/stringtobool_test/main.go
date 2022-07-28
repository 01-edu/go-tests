package main

import (
	"fmt"
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"true",
		"false",
		"TRUE",
		"FALSE",
		"1",
		"0",
		"Hello",
		"World",
		"T",
		"F",
	}
	for _, s := range args {
		challenge.Function("StringToBool", student.StringToBool, solutions.StringToBool, s)
		fmt.Println(solutions.StringToBool(s))
	}
}
