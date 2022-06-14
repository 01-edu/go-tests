package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{
		"",
		"0",
		"-551",
		"+551",
		"hekiii",
		"sda",
		"2147483647",
		"5898",
		"697abc8",
		"-2147483648"}
	for _, arg := range table {
		challenge.Function("StrisNegative", student.StrisNegative, solutions.StrisNegative, arg)
	}
}
