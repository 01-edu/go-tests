package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}
	args = append(args, random.MultRandWords()...)

	for _, arg := range args {
		challenge.Program("rostring", arg)
	}
	challenge.Program("rostring")
	challenge.Program("rostring", "this", "is")
	challenge.Program("rostring", "not", "good", "for  you")
}
