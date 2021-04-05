package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		s      string
		toFind string
	}

	table := []node{}

	// the first 15 values are valid for this test
	for i := 0; i < 15; i++ {
		wordToTest := random.RandASCII()
		firstLetterIndex := random.IntBetween(0, (len(wordToTest)-1)/2)
		lastLetterIndex := random.IntBetween(firstLetterIndex, len(wordToTest)-1)

		val := node{
			s:      wordToTest,
			toFind: wordToTest[firstLetterIndex:lastLetterIndex],
		}
		table = append(table, val)
	}

	// the next 15 values are supposed to be invalid for this test
	for i := 0; i < 15; i++ {
		wordToTest := random.RandASCII()
		wrongMatch := random.RandASCII()

		val := node{
			s:      wordToTest,
			toFind: wrongMatch,
		}
		table = append(table, val)
	}
	// those are the test values from the README examples
	table = append(table,
		node{s: "Hello!", toFind: "l"},
		node{s: "Salut!", toFind: "alu"},
		node{s: "Ola!", toFind: "hOl"},
	)

	for _, arg := range table {
		challenge.Function("Index", student.Index, solutions.Index, arg.s, arg.toFind)
	}
}
