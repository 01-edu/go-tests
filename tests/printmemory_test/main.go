package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [10]int{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = random.IntBetween(13, 126)
		}
		challenge.Function("PrintMemory", student.PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	challenge.Function("PrintMemory", student.PrintMemory, solutions.PrintMemory, table2)
}
