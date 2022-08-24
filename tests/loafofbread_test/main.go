package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"Loaf of bread", "amazing loaf of bread", "bread crumbles", "", "bread slice", "bread"}
	for _, s := range table {
		challenge.Function("LoafOfBread", student.LoafOfBread, solutions.LoafOfBread, s)
	}
}
