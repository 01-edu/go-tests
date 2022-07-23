package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{""},
		{"Hello", "Hello"},
		{"World", "World"},
		{"01010101010"},
		{"                     "},
		{"Z01 is the best place in the world", "isn't it?"},
	}

	args = append(args, random.StrSlice(chars.Words))

	for _, v := range args {
		challenge.Program("replaceeven", v...)
	}
}
