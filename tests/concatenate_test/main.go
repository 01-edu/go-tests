package main

import "github.com/01-edu/go-tests/lib/challenge"

func main() {
	table := [][]string{
		{},
		{"", "a", "b", "c"},
		{"a", "b", "c", "d"},
		{"Hello", "World", "!"},
		{"Hello", "World", "!"},
		{"Hello", ""},
		{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		{"hello there", "how are you?"},
		{"hello there", " "},
	}
	for _, arg := range table {
		challenge.Program("concatenate", arg...)
	}
}
