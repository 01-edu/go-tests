package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	arg := lib.MultRandInt()
	arg = append(arg, lib.RandInt())
	for i := 0; i < 15; i++ {
		lib.Function("Abort", student.Abort, solutions.Abort, arg[0], arg[1], arg[2], arg[3], arg[4])
		arg = lib.MultRandInt()
		arg = append(arg, lib.RandInt())
	}
}
