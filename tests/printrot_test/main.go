package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := []string{
		"A",
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
		"Z",
		"2 arguments",
		"4 arguments so invalid",
	}
	for _, s := range table {
		challenge.Program("printrot", strings.Fields(s)...)
	}
}
