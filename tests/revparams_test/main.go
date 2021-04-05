package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	challenge.Program("revparams", "choumi", "is", "the", "best", "cat")
	challenge.Program("revparams", random.StrSlice(chars.Words)...)
}
