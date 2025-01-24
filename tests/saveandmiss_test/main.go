package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	structs := []struct {
		str string
		itt int
	}{
		{"123456789", 1},
		{"e 5S @ 8* 7 =56 ;", 2},
		{"QKplq%QSw", 3},
		{"", 4},
		{"hello \\! n4ght cr3a8ure7 ", 5},
		{"Kimetsu no Yaiba", 6},
		{"8595485-52", 7},
		{"abcdefghijklmnopqrstuvwyz", 8},
		{"w58tw7474abc", 9},
		{"Po65 4o", 10},
	}

	for _, v := range structs {
		challenge.Function("SaveAndMiss", student.SaveAndMiss, solutions.SaveAndMiss, v.str, v.itt)
	}
}
