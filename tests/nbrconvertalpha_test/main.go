package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	table := []string{
		"--upper 8 5 25",
		"8 5 12 12 15",
		"12 5 7 5 14 56 4 1 18 25",
		"12 5 7 5 14 56 4 1 18 25 32 86 h",
		"32 86 h",
		"",
	}
	for i := 0; i < 5; i++ {
		m := random.IntSliceBetween(1, 46)
		s := ""
		for _, j := range m {
			s += strconv.Itoa(j) + " "
		}
		table = append(table, s)
	}

	for _, args := range table {
		challenge.Program("nbrconvertalpha", strings.Fields(args)...)
	}
}
