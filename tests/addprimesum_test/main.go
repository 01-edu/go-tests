package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/is"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	// adds random numbers
	table := rand.IntsBetween(1, 10000)

	// fill with all prime numbers between 0 and 100
	for i := 0; i < 100; i++ {
		if is.Prime(i) {
			table = append(table, i)
		}
	}

	for _, i := range table {
		challenge.Program("addprimesum", strconv.Itoa(i))
	}
	// special cases
	challenge.Program("addprimesum")
	challenge.Program("addprimesum", `""`)
	challenge.Program("addprimesum", "1", "2")
}
