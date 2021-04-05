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
	// 15 unvalid strings in the table
	table := random.StrSlice(chars.ASCII)

	// 15 valid lowercase strings of random size between 1 and 20 letters in the table
	for i := 0; i < 15; i++ {
		letters := []rune("\a\b\f\r\n\v\t")
		size := random.IntBetween(1, 20)
		r := make([]rune, size)
		for i := range r {
			r[i] = letters[rand.Intn(len(letters))]
		}
		table = append(table, string(r))
	}

	// Special cases added to table
	table = append(table,
		"Hello! How are you?",
		"a",
		"z",
		"!",
		"HelloHowareyou",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!",
		"0123456789",
		"01,02,03",
		"abcdefghijklmnopqrstuvwxyz",
		"hello",
		"hello!",
		"hello\n",
		"\n",
	)
	for _, arg := range table {
		challenge.Function("IsPrintable", student.IsPrintable, solutions.IsPrintable, arg)
	}
}
