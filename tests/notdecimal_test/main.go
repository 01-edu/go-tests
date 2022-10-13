package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{"-19.525856", "00.02", "56s44", "", "415.458", "1.6", "165", "502.3254", "51.3+95.9", "-5.0f00d00", "-53.124"}

	for i := 0; i < len(args); i++ {
		challenge.Function("NotDecimal", student.NotDecimal, solutions.NotDecimal, args[i])
	}
}
