package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"2", "5", "u", "19"},
		{"2"},
		{"1", "2", "3", "5", "7", "24"},
		{"6", "12", "24"},
		{""},
		{"-10", "10", "20", "30", "40", "50"},
		{"Hello world"},
		{"10", "20", "-1"},
		{"1", "-2", "3"},
	}
	for _, v := range args {
		challenge.Program("paramrange", v...)
	}
	challenge.Program("paramrange")
}
