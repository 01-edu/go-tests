package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type nodeTest struct {
		data []string
	}

	table := []nodeTest{}

	example := nodeTest{
		data: []string{"Hello! How are You?"},
	}
	table = append(table, example)

	for i := 0; i < 5; i++ {
		val := nodeTest{
			data: random.StrSlice(chars.Words),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		for _, s := range arg.data {
			challenge.Function("Rot14", student.Rot14, solutions.Rot14, s)
		}
	}
}
