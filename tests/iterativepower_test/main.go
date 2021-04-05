package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	i := 0
	for i < 30 {
		nb := lib.RandIntBetween(-8, 8)
		power := lib.RandIntBetween(-10, 10)
		lib.Function("IterativePower", student.IterativePower, solutions.IterativePower, nb, power)
		i++
	}
	lib.Function("IterativePower", student.IterativePower, solutions.IterativePower, 0, 0)
	lib.Function("IterativePower", student.IterativePower, solutions.IterativePower, 0, 1)
}
