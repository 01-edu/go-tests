package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"First78last", "     ", " 280jsl", "he", "", "honey!", "Z", "email123@live.fr", "w45m$", "52", "474abc", "0"}
	for _, s := range table {
		challenge.Function("PrintIf", student.PrintIf, solutions.PrintIf, s)
	}
}
