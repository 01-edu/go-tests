package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"hello", "", "Kimetsu no Yaiba", "Z", "123@live.fr", "write %d ==> 45m$", "-552"}
	for _, s := range table {
		challenge.Function("RetainFirstHalf", student.RetainFirstHalf, solutions.RetainFirstHalf, s)
	}
}
