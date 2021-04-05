package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []int{lib.RandIntBetween(2, 20)}
	args = append(args, lib.MultRandIntBetween(2, 20)...)
	args = append(args, lib.MultRandIntBetween(2, 20)...)

	for _, v := range args {
		lib.Function("ActiveBits", student.ActiveBits, solutions.ActiveBits, v)
	}
}
