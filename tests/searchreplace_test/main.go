package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	type nodeTest struct {
		dataSearched    string
		letterLookedFor string
		letterReplacing string
	}

	args := []nodeTest{
		{
			dataSearched:    "hélla",
			letterLookedFor: "é",
			letterReplacing: "o",
		},
		{
			dataSearched:    "hella",
			letterLookedFor: "z",
			letterReplacing: "o",
		},
		{
			dataSearched:    "hella",
			letterLookedFor: "h",
			letterReplacing: "o",
		},
	}

	for i := 0; i < 20; i++ {
		args = append(args, nodeTest{
			dataSearched:    random.Str(chars.Words, 13),
			letterLookedFor: random.Str(chars.Alnum, 1),
			letterReplacing: random.Str(chars.Alnum, 1),
		})
	}

	for _, arg := range args {
		challenge.Program("searchreplace", arg.dataSearched, arg.letterLookedFor, arg.letterReplacing)
	}
}
