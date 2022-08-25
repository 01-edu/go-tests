package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"1010101010", "plq%QSw", "", "Life Is Great", "Z", "johndoe@imposssible.io", "8595485-52", "-552", "w58tw7474abc", "Mine Craft"}
	for _, s := range table {
		challenge.Function("JumpOver", student.JumpOver, solutions.JumpOver, s)
	}
}
