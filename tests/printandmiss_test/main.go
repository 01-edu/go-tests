package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"1234556789", "e 5£ @ 8* 7 =56 ;", "QKplq%QSw", "", "hello \\! n4ght cr3a8ure7 ", "Kimetsu no Yaiba", "8595485-52", "abcdefghijklmnopqrstuvwyz", "w58tw7474abc", "Po65 4o"}
	_itt := []int{1, 2, 3, 4, 5, 0, -25, 8, -1, 8}
	i := 0
	j := 0
	for i < len(table) && j < len(_itt) {
		challenge.Function("PrintAndMiss", student.PrintAndMiss, solutions.PrintAndMiss, table[i], _itt[j])
		i++
		j++
	}
}
