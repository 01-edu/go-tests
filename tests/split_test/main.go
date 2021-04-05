package main

import (
	"math/rand"
	"strings"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	separators := []string{
		"!=HA=!",
		"!==!",
		"    ",
		"|=choumi=|",
		"|<=>|",
		random.Str(chars.Upper, 3),
		"<<==123==>>",
		"[<>abc<>]",
	}

	type node struct {
		s   string
		sep string
	}
	args := []node{{s: "HelloHAhowHAareHAyou?", sep: "HA"}}
	for i := 0; i < 15; i++ {
		separator := separators[rand.Intn(len(separators))]
		args = append(args, node{
			s:   strings.Join(random.StrSlice(chars.Alnum), separator),
			sep: separator,
		})
	}
	for _, arg := range args {
		challenge.Function("Split", student.Split, solutions.Split, arg.s, arg.sep)
	}
}
