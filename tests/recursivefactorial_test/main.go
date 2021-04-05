package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := append(
		lib.MultRandInt(),
		lib.IntRange(0, 20)...,
	)
	for _, arg := range table {
		lib.Function("RecursiveFactorial", student.RecursiveFactorial, solutions.RecursiveFactorial, arg)
	}
}
