package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	var (
		flags          = []string{"--insert=", "--order"}
		flagsShorthand = []string{"-i=", "-o"}
		randArgFlag    = []string{random.Str(chars.Words, 13), random.Str(chars.Words, 13)}
		randArg        = []string{random.Str(chars.Words, 13), random.Str(chars.Words, 13), ""}
	)

	challenge.Program("flags", flagsShorthand[0]+"v2", "v1")
	challenge.Program("flags", flagsShorthand[1], "v1")
	challenge.Program("flags", "-h")
	challenge.Program("flags", "--help")
	challenge.Program("flags")

	for _, v2 := range randArgFlag {
		for _, v1 := range randArg {
			challenge.Program("flags", flags[0]+v2, flags[1], v1)
		}
	}
	for _, v2 := range randArgFlag {
		for _, v1 := range randArg {
			challenge.Program("flags", flagsShorthand[0]+v2, flagsShorthand[1], v1)
		}
	}
}
