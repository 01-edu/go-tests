package main

import (
	"github.com/01-edu/go-tests/lib"
)

func main() {
	lib.Program("revparams", "choumi", "is", "the", "best", "cat")
	lib.Program("revparams", lib.MultRandWords()...)
}
