package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	args := []string{
		"FOR PONY",
		"this ... is sparta, then again, maybe not",
		" ",
		" lorem,ipsum ",
	}

	args = append(args, rand.MultRandWords()...)

	for _, v := range args {
		challenge.Program("lastword", v)
	}

	challenge.Program("lastword", "a", "b")
	challenge.Program("lastword")
}
