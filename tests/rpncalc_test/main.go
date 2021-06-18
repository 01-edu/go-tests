package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	ops := []string{"+", "-", "/", "*", "%"}
	args := []string{
		"1",
		"1 2 * 3 * 4 +",
		"3 1 2 * * 4 %",
		"5 10 9 / - 50 *",
		"32   / 22",
		"88 67 dks -",
		"     1      3 * 2 -",
	}

	for i := 0; i < 6; i++ {
		str := ""
		for j := 0; j < random.IntBetween(3, 10); j++ {
			if j%2 == 0 && j != 0 {
				str += ops[random.IntBetween(0, len(ops)-1)] + " "
			} else {
				str += strconv.Itoa(random.IntBetween(1, 100)) + " "
			}
		}
		args = append(args, str)
	}

	for _, v := range args {
		challenge.Program("rpncalc", v)
	}
	challenge.Program("rpncalc")
	challenge.Program("rpncalc", "1 2 * 3 * 4 +", "10 33 - 12 %")
}
