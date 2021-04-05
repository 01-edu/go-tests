package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		s      string
		toFind string
	}

	args := []node{
		{s: "Hello!", toFind: "l"},
		{s: "Salut!", toFind: "alu"},
		{s: "Ola!", toFind: "hOl"},
	}

	// the first 15 values are valid for this test
	for i := 0; i < 15; i++ {
		wordToTest := random.Str(chars.ASCII, 13)
		firstLetterIndex := random.IntBetween(0, (len(wordToTest)-1)/2)
		lastLetterIndex := random.IntBetween(firstLetterIndex, len(wordToTest)-1)

		val := node{
			s:      wordToTest,
			toFind: wordToTest[firstLetterIndex:lastLetterIndex],
		}
		args = append(args, val)
	}

	// the next 15 values are supposed to be invalid for this test
	for i := 0; i < 15; i++ {
		val := node{
			s:      random.Str(chars.ASCII, 13),
			toFind: random.Str(chars.ASCII, 13),
		}
		args = append(args, val)
	}

	for _, arg := range args {
		challenge.Function("Index", student.Index, solutions.Index, arg.s, arg.toFind)
	}
}
