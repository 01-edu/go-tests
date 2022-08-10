package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"123",
		"H1ll0",
		"",
		"1",
		"1.1",
		"Containe1number",
		"     ",
		"upson lorem ipsum",
	}

	for _, arg := range args {
		challenge.Function("CheckNumber", student.CheckNumber, solutions.CheckNumber, arg)
	}
}
