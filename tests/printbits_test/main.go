package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"",
		"a",
		"bc",
		"def",
		"notanumber",
		random.Str(chars.Basic, 13),
	}
	for i := 0; i < 20; i++ {
		args = append(args, strconv.Itoa(random.IntBetween(0, 255)))
	}
	for _, v := range args {
		challenge.Program("printbits", strings.Fields(v)...)
	}
}
