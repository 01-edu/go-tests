package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	years := []int{
		2016,
		2017,
		2000,
		1900,
		1920,
		1921,
		1,
		-10,
		10000,
		10001,
		10002,
		-2012,
		0,
	}
	for _, arg := range years {
		challenge.Function("LeapYear", student.LeapYear, solutions.LeapYear, arg)
	}
}
