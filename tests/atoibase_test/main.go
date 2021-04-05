package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/base"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	type node struct {
		s    string
		base string
	}

	args := []node{
		{s: "125", base: "0123456789"},
		{s: "1111101", base: "01"},
		{s: "7D", base: "0123456789ABCDEF"},
		{s: "uoi", base: "choumi"},
		{s: "bbbbbab", base: "-ab"},
	}

	// 15 random pairs of string numbers with valid bases
	for i := 0; i < 15; i++ {
		validBaseToInput := base.Valid()
		val := node{
			s:    base.StringFrom(validBaseToInput),
			base: validBaseToInput,
		}
		args = append(args, val)
	}
	// 15 random pairs of string numbers with invalid bases
	for i := 0; i < 15; i++ {
		invalidBaseToInput := base.Invalid()
		val := node{
			s:    "thisinputshouldnotmatter",
			base: invalidBaseToInput,
		}
		args = append(args, val)
	}
	for _, arg := range args {
		challenge.Function("AtoiBase", student.AtoiBase, solutions.AtoiBase, arg.s, arg.base)
	}
}
