package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"student", "hello Iyan", "world", "Hello! How are you?"}
	for _, arg := range table {
		challenge.Function("VowelsIndex", student.vowelsindex, solutions.VowelsIndex, arg)
	}
}
