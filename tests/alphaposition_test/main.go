package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []rune{
		'a',
		'A',
		'z',
		'Z',
		'0',
		'9',
		'!',
		'B',
		'@',
		'b',
		'c',
		'd',
	}

	for _, arg := range args {
		challenge.Function("AlphaPosition", student.AlphaPosition, solutions.AlphaPosition, arg)
	}
}
