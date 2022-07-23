package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"ReverseBits",
		"Hello, world!",
		"My name is John",
		"I am a student",
		"Five is 5",
		"Seven never sevens",
		"Eight is 8",
		"H",
		"I",
	}

	for _, v := range args {
		challenge.Function("RemoveOdd", student.RemoveOdd,solutions.RemoveOdd, v)
	}
}
