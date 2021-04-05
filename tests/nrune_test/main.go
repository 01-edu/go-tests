package main

import (
	"math/rand"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		word string
		n    int
	}

	args := []node{
		{word: "Hello!", n: 3},
		{word: "Salut!", n: 2},
		{word: "Ola!", n: 4},
		{word: "♥01!", n: 1},
		{word: "Not", n: 6},
		{word: "Also not", n: 9},
		{word: "Tree house", n: 5},
		{word: "Whatever", n: 0},
		{word: "Whatshisname", n: -2},
	}

	for i := 0; i < 30; i++ {
		wordToInput := random.Str(chars.ASCII, 13)
		args = append(args, node{
			word: wordToInput,
			n:    rand.Intn(len(wordToInput)) + 1,
		})
	}

	for _, arg := range args {
		challenge.Function("NRune", student.NRune, solutions.NRune, arg.word, arg.n)
	}
}
