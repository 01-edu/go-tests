package main

import "github.com/01-edu/go-tests/lib/challenge"

func main() {
	table := [][]string{
		{},
		{"a"},
		{"", "a", "b", "c"},
		{"a", "b", "c", "d"},
		{"Hello", "World", "!"},
		{"Hello", "World", "!"},
		{"Hello", ""},
		{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
	}
	for _, arg := range table {
		challenge.Program("printevenarguments", arg...)
	}
}
