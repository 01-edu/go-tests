package main

import (
	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	challenge.Program("revparams", "choumi", "is", "the", "best", "cat")
	challenge.Program("revparams", lib.MultRandWords()...)
}
