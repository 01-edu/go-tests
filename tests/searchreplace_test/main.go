package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	type nodeTest struct {
		dataSearched    string
		letterLookedFor string
		letterReplacing string
	}

	table := []nodeTest{}

	for i := 0; i < 20; i++ {
		letter1 := []rune(random.RandAlnum())
		letter2 := []rune(random.RandAlnum())

		table = append(table,
			nodeTest{
				dataSearched:    random.RandWords(),
				letterLookedFor: string(letter1[0]),
				letterReplacing: string(letter2[0]),
			})
	}

	table = append(table,
		nodeTest{
			dataSearched:    "hélla",
			letterLookedFor: "é",
			letterReplacing: "o",
		},
		nodeTest{
			dataSearched:    "hella",
			letterLookedFor: "z",
			letterReplacing: "o",
		},
		nodeTest{
			dataSearched:    "hella",
			letterLookedFor: "h",
			letterReplacing: "o",
		},
	)

	for _, arg := range table {
		challenge.Program("searchreplace", arg.dataSearched, arg.letterLookedFor, arg.letterReplacing)
	}
}
