package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]string{{"Hello!", " How are you?"}}

	// 30 valid pair of ramdom strings to concatenate
	for i := 0; i < 30; i++ {
		args = append(args, []string{random.Str(chars.ASCII, 13), random.Str(chars.ASCII, 13)})
	}
	for _, arg := range args {
		challenge.Function("Concat", student.Concat, solutions.Concat, arg[0], arg[1])
	}
}
