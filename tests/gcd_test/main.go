package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{"42", "10"},
		{"42", "12"},
		{"14", "77"},
		{"17", "3"},
		{"23"},
		{"12", "23"},
		{"25", "15"},
		{"23043", "122"},
		{"11", "77"},
	}
	for i := 0; i < 5; i++ {
		a := strconv.Itoa(random.IntBetween(1, 100000))
		b := strconv.Itoa(random.IntBetween(1, 100))
		args = append(args, []string{a, b})
	}
	for _, v := range args {
		challenge.Program("gcd", v...)
	}
}
