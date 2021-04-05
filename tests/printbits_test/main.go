package main

import (
	"strconv"
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	var args []string
	for i := 0; i < 20; i++ {
		args = append(args, strconv.Itoa(rand.IntBetween(0, 255)))
	}
	args = append(args,
		"",
		"a",
		"bc",
		"def",
		"notanumber",
		rand.RandBasic(),
	)
	for _, v := range args {
		challenge.Program("printbits", strings.Fields(v)...)
	}
}
