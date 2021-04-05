package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	args = append(args, lib.MultRandWords()...)

	for _, v := range args {
		lib.Program("lastword", v)
	}

	lib.Program("lastword", "a", "b")
	lib.Program("lastword")
}
