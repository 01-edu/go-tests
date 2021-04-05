package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	challenge.Program("revparams", "choumi", "is", "the", "best", "cat")
	challenge.Program("revparams", rand.MultRandWords()...)
}
