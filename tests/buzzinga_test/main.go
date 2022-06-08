package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
)

func solutions(n int32) string {
	solutions := map[int32]string{
		15: "*\n*\nBuz\n*\nzinga\nBuz\n*\n*\nBuz\nzinga\n*\nBuz\n*\n*\nBuzzinga\n",
		10: "*\n*\nBuz\n*\nzinga\nBuz\n*\n*\nBuz\nzinga\n",
		5:  "*\n*\nBuz\n*\nzinga\n",
		4:  "*\n*\nBuz\n*\n",
		20: "*\n*\nBuz\n*\nzinga\nBuz\n*\n*\nBuz\nzinga\n*\nBuz\n*\n*\nBuzzinga\n*\n*\nBuz\n*\nzinga\n",
	}
	return solutions[n]
}

func main() {
	challenge.Function("BuzZinga", student.BuzZinga, solutions, 20)
	challenge.Function("BuzZinga", student.BuzZinga, solutions, 15)
	challenge.Function("BuzZinga", student.BuzZinga, solutions, 10)
	challenge.Function("BuzZinga", student.BuzZinga, solutions, 5)
	challenge.Function("BuzZinga", student.BuzZinga, solutions, 4)
}
