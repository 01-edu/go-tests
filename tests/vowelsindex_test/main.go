package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"student", "hello Iyan", "world","ptghjkvl", "Hello! How are you?", "AAEO$%&UOTX", "", "AEMNOoa*+-aeiou", "wo5opi=ws8;eza"}

	for _, arg := range table {
		challenge.Function("VowelsIndex", student.VowelsIndex, solutions.VowelsIndex, arg)
	}
}
