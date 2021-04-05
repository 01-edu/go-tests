package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	for i := 0; i < 10; i++ {
		start := lib.RandIntBetween(-20, 20)
		end := lib.RandIntBetween(-20, 20)
		challenge.Program("reverserange", strconv.Itoa(start), strconv.Itoa(end))
	}
	challenge.Program("reverserange", "2", "1", "3")
	challenge.Program("reverserange", "a", "1")
	challenge.Program("reverserange", "1", "b")
	challenge.Program("reverserange", "1", "nan")
	challenge.Program("reverserange", "nan", "b")
	challenge.Program("reverserange")
}
