package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	args := []string{
		"",
		"Hello",
		"World",
		"Home",
		"Theorem",
		"Choumi is the best cat",
		"abracadaba 01!",
		"abc",
		"MaTheMatiCs",
	}

	args = append(args, lib.MultRandAlnum()...)

	for _, v := range args {
		lib.Program("repeatalpha", v)
	}
	lib.Program("repeatalpha")
	lib.Program("repeatalpha", "", "")
}
