package main

import (
	"math/rand"

	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	for i := 0; i < 20; i++ {
		token := rand.Perm(20)
		target := rand.Intn(30)
		lib.Function("TwoSum", student.TwoSum, solutions.TwoSum, token, target)
	}
}
