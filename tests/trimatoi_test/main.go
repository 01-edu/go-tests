package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func stringsToTrimAtoi(a []string) []string {
	for i := 0; i < 4; i++ {
		n := random.Str(chars.Digit, random.IntBetween(1, 2))
		s := random.Str(chars.Alnum, random.IntBetween(0, 3))
		rand := random.IntBetween(0, 14)
		if rand <= 4 {
			n += "-"
			s += "-"
		}
		s += random.Str(chars.Alnum, random.IntBetween(0, 10))
		a = append(a, s, n)
	}
	return a
}

func main() {
	args := stringsToTrimAtoi([]string{
		"",
		"12345",
		"str123ing45",
		"012 345",
		"Hello World!",
		"sd+x1fa2W3s4",
		"sd-x1fa2W3s4",
		"sdx1-fa2W3s4",
		"123-4asd",
		random.Str(chars.Alnum, 13),
	})
	for _, arg := range args {
		challenge.Function("TrimAtoi", student.TrimAtoi, solutions.TrimAtoi, arg)
	}
}
