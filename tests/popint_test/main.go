package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	
	args := [][]int{
		{1, 3, 5},
		{-1, -2, 3},
		{},
		{22134, 2345345, 575672, 2567567, 2456456, 2345345},
	}

	for _, arg := range args {
		challenge.Function("PopInt", student.PopInt, solutions.PopInt, arg)
	}
}
