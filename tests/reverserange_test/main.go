package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := [][]string{
		{"2", "1", "3"},
		{"-1", "2"},
		{"0", "-3"},
		{"1", "nan"},
		{"nan", "b"},
	}
	for i := 0; i < 10; i++ {
		start := random.IntBetween(-20, 20)
		end := random.IntBetween(-20, 20)
		table = append(table, []string{strconv.Itoa(start), strconv.Itoa(end)})
	}
	for _, v := range table {
		challenge.Program("reverserange", v...)
	}
}
