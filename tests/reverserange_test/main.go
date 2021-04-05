package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	for i := 0; i < 10; i++ {
		start := lib.RandIntBetween(-20, 20)
		end := lib.RandIntBetween(-20, 20)
		lib.Program("reverserange", strconv.Itoa(start), strconv.Itoa(end))
	}
	lib.Program("reverserange", "2", "1", "3")
	lib.Program("reverserange", "a", "1")
	lib.Program("reverserange", "1", "b")
	lib.Program("reverserange", "1", "nan")
	lib.Program("reverserange", "nan", "b")
	lib.Program("reverserange")
}
