package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"W"},
		{""},
		{"H", "world"},
		{"Hello World"},
		{"A"},
		{"B"},
		{"C"},
		{"&"},
		{"20"},
		{"100"},
		{" "},
		{"4"}, {"E"},
		{"	"},
		{"13 233 4 23"},
		{"1"},
		{"7"},
		{"8"},
		{"£"},
		{"@"},
		{"$"},
		{"~"},
		{"0"},
		{"*"},
	}
	for _, s := range args {
		challenge.Program("getascii", s...)
	}
}
