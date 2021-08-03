package main

import (
	"strings"
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

type functionArgs struct {
	Text string
	N    int
}

func main() {
	args := []functionArgs{
		{Text: "Hi Hi Hello", N: 1},
		{Text: "Orange Orange is the sun sliding to the horizon after a summer day. Orange is the sound of dribbling basetball. Orange is the smell of a tiger lily petal. Orange is the taste of thirst-quenching Nehi Soda. Orange is the color of peach marmalade on a side of toast. Orange is the sound of a carrot popping out of the ground.", N: 3},
		{Text: "toto toto toto", N: 1},
		{Text: "tutu tutu tutu tata tata toto apo", N: 3},
	}

	for i := 0; i < 30; i++ {
		wordToInput := random.Str("x ", 100)
		set := make(map[string]bool) // New empty set
		for _, k := range strings.Split(wordToInput, " ") {
			set[k] = true
		}

		args = append(args, functionArgs{
			Text: wordToInput,
			// The set is used to make sure we are asking for an index
			// that is in range of the number of keys in the map.
			N: random.IntBetween(1, len(set)),
		})
	}

	for _, arg := range args {
		challenge.Function("MaxWordCountN", student.MaxWordCountN, solutions.MaxWordCountN, arg.Text, arg.N)
	}
}
