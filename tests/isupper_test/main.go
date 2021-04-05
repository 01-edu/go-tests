package main

import (
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
		size := random.IntBetween(1, 20)
		randLow := random.Str(chars.Lower, 13)
		if len(randLow) <= size {
			table = append(table, randLow)
		} else {
			table = append(table, randLow[:size])
		}
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
		"HELLO",
		"HELLO!",
	)
	for _, arg := range table {
		challenge.Function("IsUpper", student.IsUpper, solutions.IsUpper, arg)
	}
}
