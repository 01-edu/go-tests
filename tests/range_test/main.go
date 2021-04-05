package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	for i := 0; i < 10; i++ {
		start := lib.RandIntBetween(-20, 20)
		end := lib.RandIntBetween(-20, 20)
		lib.Program("range", strconv.Itoa(start), strconv.Itoa(end))
	}
	lib.Program("range", "2", "1", "3")
	lib.Program("range", "a", "1")
	lib.Program("range", "1", "b")
	lib.Program("range", "1", "nan")
	lib.Program("range", "nan", "b")
	lib.Program("range")
}
