package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	lib.ChallengeMain("revparams", "choumi", "is", "the", "best", "cat")
	lib.ChallengeMain("revparams", lib.MultRandWords()...)
}
