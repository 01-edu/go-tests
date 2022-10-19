package main

import (
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
		"t",
		"True",
	}
	for _, s := range args {
		challenge.Function("StringToBool", student.StringToBool, solutions.StringToBool, s)
	}
}
