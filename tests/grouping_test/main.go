package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += random.Str(chars.Lower, 1)
		if random.Int()%2 == 0 {
			result += random.Str(chars.Lower, 1)
		}
		if i != n-1 {
			result += "|"
		}
	}

	result += ")"
	return result
}

func main() {
	args := [][2]string{
		{"(a)", "I'm heavyjumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"},
		{"(e|n)", "I currently have 4 windows opened up… and I don’t know why."},
		{"(hi)", "He swore he just saw his sushi move."},
		{"(s)", ""},
		{"i", "Something in the air"},
		{validRegExp(2), random.Str(chars.Words, 60)},
		{validRegExp(2), random.Str(chars.Words, 60)},
		{validRegExp(6), random.Str(chars.Words, 60)},
		{random.Str("axyz", 1), random.Str("axyzdassbzzxxxyy cdq     ", 10)},
	}
	for _, s := range args {
		challenge.Program("grouping", s[0], s[1])
	}
}
