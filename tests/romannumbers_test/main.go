package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		"0",
		"4000",
		"5000",
		"12433",
		"hello",
		"good luck",
	}
	for i := 0; i < 7; i++ {
		table = append(table, strconv.Itoa(random.IntBetween(0, 4000)))
	}
	for _, v := range table {
		challenge.Program("romannumbers", v)
	}
}
