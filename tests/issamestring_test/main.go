package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]string{
		{"abcd", "AbCd"},
		{"abcd", "abcddd"},
		{"ab12", "abcd"},
		{"", "abcd"},
		{"ABCD", "abcd"},
		{"abcd12HJK", "abcd12HJK"},
		{"+++\t----", "+++\t----"},
		{"abcdefgh", "abcdefghijklmnopqrstuvwxyz"},
		{"	", "		"},
		{"   ", "   "},
	}
	for _, arg := range args {
		challenge.Function("IsSameString", student.IsSameString, solutions.IsSameString, arg[0], arg[1])
	}
}
