package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {

	args := [][]string{
		{"John Doe"},
		{"Hello World"},
		{""},
		{" "},
		{"   Hello World"},
		{"Hello World   "},
		{"Hello World   My Name Is John Doe"},
		{"Peter Parker!@#$%^&*()_+"},
		{"James Bond\n"},
		{"Nick Fur y"},
		{"hello", "world"},
		{"hello", "world", "nick"},
	}
	for _, arg := range args {
		challenge.Program("swapname", arg...)
	}
}
