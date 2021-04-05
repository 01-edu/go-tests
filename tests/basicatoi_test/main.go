package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(lib.RandPosZ())
	}
	table = append(table,
		strconv.Itoa(lib.MaxInt),
		"",
		"0",
		"12345",
		"0000012345",
		"000000",
	)
	for _, arg := range table {
		lib.Function("BasicAtoi", student.BasicAtoi, solutions.BasicAtoi, arg)
	}
}
