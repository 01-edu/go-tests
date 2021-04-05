package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := [][]string{
		{"0", "0"},
		{"1", "2"},
		{"2"},
		{"0"},
		{"3", "3"},
		{"1", "5"},
		{"5", "1"},
		{"4", "3"},
	}
	for i := 0; i < 2; i++ {
		table = append(table, []string{
			strconv.Itoa(random.IntBetween(1, 200)),
			strconv.Itoa(random.IntBetween(1, 200)),
		})
	}

	for _, v := range table {
		challenge.Program("printchessboard", v...)
	}
}
