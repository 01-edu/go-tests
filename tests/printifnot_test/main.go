package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"ab", "		", "good morning sunshine!", " 28", "he", "", "honey!", "Z", "email123@live.fr", "w45m$", "-552", "474abc", "<="}
	for _, s := range table {
		challenge.Function("PrintIfNot", student.PrintIfNot, solutions.PrintIfNot, s)
	}
}
