package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{""},
		{"0"},
		{"3", "5"},
		{"5"},
		{"9"},
		{"24"},
		{"99"},
		{"100"},
	}
	for _, arg := range args {
		challenge.Program("pingpong", arg...)
	}
}
