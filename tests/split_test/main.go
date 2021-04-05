package main

import (
	"math/rand"
	"strings"

	"github.com/01-edu/go-tests/lib/random"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	separators := []string{
		"!=HA=!",
		"!==!",
		"    ",
		"|=choumi=|",
		"|<=>|",
		random.RandStr(3, random.Upper),
		"<<==123==>>",
		"[<>abc<>]",
	}

	type node struct {
		s   string
		sep string
	}
	table := []node{}
	// 15 random slice of strings

	for i := 0; i < 15; i++ {
		separator := separators[rand.Intn(len(separators))]
		val := node{
			s:   strings.Join(random.MultRandAlnum(), separator),
			sep: separator,
		}
		table = append(table, val)
	}

	table = append(table,
		node{s: "HelloHAhowHAareHAyou?", sep: "HA"})

	for _, arg := range table {
		challenge.Function("Split", student.Split, solutions.Split, arg.s, arg.sep)
	}
}
