package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/is"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	// adds random numbers
	args := random.IntSliceBetween(1, 10000)

	// fill with all prime numbers between 0 and 100
	for i := 0; i < 100; i++ {
		if is.Prime(i) {
			args = append(args, i)
		}
	}

	for _, i := range args {
		challenge.Program("addprimesum", strconv.Itoa(i))
	}
	// special cases
	challenge.Program("addprimesum")
	challenge.Program("addprimesum", `""`)
	challenge.Program("addprimesum", "1", "2")
}
