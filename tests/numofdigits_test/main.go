package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	ar := []int{
		1,
		5858,
		485,
		5,
		-525,
		47,
		85955,
		0,
		-8,
		5285,
		-48,
	}
	for i := 0; i < len(ar); i++ {
		challenge.Function("Numofdigits", student.Numofdigits, solutions.Numofdigits, ar[i])
	}
}
