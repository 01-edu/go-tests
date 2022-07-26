package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {

	args := [][]string{
		{"John Doe"},
		{"James Davud"},
		{"Mister Bie"},
		{" "},
		{"   Hello World"},
		{"Hello World   "},
		{"Hello World   My Name Is John Doe"},
		{"Peter Parker!@#$%^&*()_+"},
		{"James Bond\n"},
		{"Nick Fur y"},
		{"hello", "world"},
		{"hello", "world", "nick"},
		{" Johan      Paven "},
	}
	for _, arg := range args {
		challenge.Program("swapname", arg...)
	}
}
