package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	args := []string{
		"abc   ",
		"Let there     be light",
		"     AkjhZ zLKIJz , 23y",
		"",
	}
	args = append(args, lib.MultRandWords()...)

	for _, arg := range args {
		lib.Program("rostring", arg)
	}
	lib.Program("rostring")
	lib.Program("rostring", "this", "is")
	lib.Program("rostring", "not", "good", "for  you")
}
