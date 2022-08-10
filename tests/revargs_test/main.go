package main

import (

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"Hello"},
		{""},
		{},
		{"Hello World", "world"},
		{"Hello World", "world", "Hello World"},
		{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"},
		{"30", "20", "10", "40", "50", "60", "70", "80", "90", "100"},
		{"A B C D E"},
		{"4 3 2 1 0"},
		{"13 233 4 23"},
	}
	for _, s := range args {
		challenge.Program("revargs", s...)
	}
}
