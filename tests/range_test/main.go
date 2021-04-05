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
		challenge.Program("range", strconv.Itoa(start), strconv.Itoa(end))
	}
	challenge.Program("range", "2", "1", "3")
	challenge.Program("range", "a", "1")
	challenge.Program("range", "1", "b")
	challenge.Program("range", "1", "nan")
	challenge.Program("range", "nan", "b")
	challenge.Program("range")
}
