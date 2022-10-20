package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"4"},
		{""},
		{"52"},
		{"8"},
		{"42"},
		{"92"},
		{"102"},
		{"127"},
		{"40"},
		{"15"},
		{"9"},
		{"4"},
		{"E"},
		{"	"},
		{"13"},
		{"68"},
		{"77"},
		{"87"},
		{"51"},
		{"2"},
		{"3"},
		{"15"},
		{"16"},
		{"0"},
	}
	for _, s := range args {
		challenge.Program("getalpha", s...)
	}
}
