package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := [][]string{
		{},
		{""},
		{" "},
		{"", "a"},
		{"  "},
		{" ", "b", "c"},
		{"a", ""},
		{"first", "second"},
		{"first", "", "third"},
		{"this", "is", "a", "longer", "test", "it", "is", "not", "over", "yet", "and", "still", "continues", "and", "it", "is", "over", "now"},
		{"longswordsuperlongwithalotofcharactersdigits012341234andspecialcharacter*(())){}@€s£$^", "longswordsuperlongwithalotofcharactersdigits012341234andspecialcharacters£$^"},
	}
	for _, s := range table {
		challenge.Function("ByebyeFirst", student.ByeByeFirst, solutions.ByeByeFirst, s)
	}
}
