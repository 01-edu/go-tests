package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]string{}

	// 30 valid pair of ramdom slice of strings to concatenate
	for i := 0; i < 30; i++ {
		table = append(table, random.StrSlice(chars.ASCII))
	}
	table = append(table,
		[]string{"Hello!", " How are you?", "well and yourself?"},
	)
	for _, arg := range table {
		challenge.Function("BasicJoin", student.BasicJoin, solutions.BasicJoin, arg)
	}
}
