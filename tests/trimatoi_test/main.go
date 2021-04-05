package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func stringsToTrimAtoi(a []string) []string {
	for index := 0; index < 4; index++ {
		s := ""
		s += random.RandStr(random.IntBetween(0, 2), random.Alnum)
		x := random.IntBetween(0, 14)
		if x <= 4 {
			s += "-"
		}
		s += random.RandStr(random.IntBetween(0, 10), random.Alnum)
		a = append(a, s)
	}
	return a
}

func main() {
	a := []string{
		"",
		"12345",
		"str123ing45",
		"012 345",
		"Hello World!",
		"sd+x1fa2W3s4",
		"sd-x1fa2W3s4",
		"sdx1-fa2W3s4",
		random.RandAlnum(),
	}
	a = stringsToTrimAtoi(a)
	for _, elem := range a {
		challenge.Function("TrimAtoi", student.TrimAtoi, solutions.TrimAtoi, elem)
	}
}

// TODO: refactor, including the subject
