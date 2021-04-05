package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += rand.RandStr(1, rand.Lower)
		if rand.Int()%2 == 0 {
			result += rand.RandStr(1, rand.Lower)
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
		{validRegExp(2), rand.RandStr(60, rand.Lower+rand.Space)},
		{validRegExp(2), rand.RandStr(60, rand.Lower+rand.Space)},
		{validRegExp(6), rand.RandStr(60, rand.Lower+rand.Space)},
		{rand.RandStr(1, "axyz"), rand.RandStr(10, "axyzdassbzzxxxyy cdq     ")},
	}
	for _, s := range args {
		challenge.Program("grouping", s[0], s[1])
	}
}
