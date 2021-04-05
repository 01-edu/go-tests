package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += random.RandStr(1, random.Lower)
		if random.Int()%2 == 0 {
			result += random.RandStr(1, random.Lower)
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
		{validRegExp(2), random.RandStr(60, random.Lower+random.Space)},
		{validRegExp(2), random.RandStr(60, random.Lower+random.Space)},
		{validRegExp(6), random.RandStr(60, random.Lower+random.Space)},
		{random.RandStr(1, "axyz"), random.RandStr(10, "axyzdassbzzxxxyy cdq     ")},
	}
	for _, s := range args {
		challenge.Program("grouping", s[0], s[1])
	}
}
