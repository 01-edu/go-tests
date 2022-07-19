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
		"585 41 -9",
		"hekiii",
		"sda",
		"2147483647",
		"5898",
		"697abc8",
		"82-58",
		"-2147483648",
		" -585",
		"10 15",
		"975 ",
	}
	for _, arg := range table {
		challenge.Function("StrisNegative", student.StrisNegative, solutions.StrisNegative, arg)
	}
}
