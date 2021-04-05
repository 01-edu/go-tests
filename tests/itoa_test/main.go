package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 50; i++ {
		challenge.Function("Itoa", student.Itoa, solutions.Itoa, random.Int())
	}
}
