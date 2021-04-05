package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][2]string{
		{"listen", "silent"},
		{"alem", "school"},
		{"neat", "a net"},
		{"anna madrigal", "a man and a girl"},
		{"abcc", "abcd"},
		{"aaaac", "caaaa"},
		{"", ""},
		{"       ", ""},
		{"lyam", "meow"},
		{"golang", "lang go"},
		{"verylongword", "v e r y l o n g w o r d"},
		{"chess", "ches"},
		{"anagram", "nnagram"},
		{"chess", "board"},
		{"mmm", "m"},
		{"pulp", "fiction"},
	}
	for i := 0; i < 15; i++ {
		args = append(args, [2]string{
			random.Str(chars.Words, random.IntBetween(15, 20)),
			random.Str(chars.Words, random.IntBetween(15, 20)),
		})
	}

	for _, arg := range args {
		challenge.Function("IsAnagram", student.IsAnagram, solutions.IsAnagram, arg[0], arg[1])
	}
}
