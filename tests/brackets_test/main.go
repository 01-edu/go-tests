package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	oneArgs := []string{
		"(johndoe)",
		")()",
		"([)]",
		"{2*[d - 3]/(12)}",
	}

	// 18 random tests ( at least half are valid)
	for i := 0; i < 3; i++ {
		oneArgs = append(oneArgs,
			"("+random.RandASCII()+")",
			"["+random.RandASCII()+"]",
			"{"+random.RandASCII()+"}",
			"("+random.RandAlnum()+")",
			"["+random.RandAlnum()+"]",
			"{"+random.RandAlnum()+"}",
		)
	}

	challenge.Program("brackets")

	for _, v := range oneArgs {
		challenge.Program("brackets", v)
	}

	multArg := [][]string{
		{"", "{[(0 + 0)(1 + 1)](3*(-1)){()}}"},
		{"{][]}", "{3*[21/(12+ 23)]}"},
		{"{([)])}", "{{{something }- [something]}}", "there are"},
	}

	for _, v := range multArg {
		challenge.Program("brackets", v...)
	}
}
