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
		table[i] = strconv.Itoa(lib.RandInt())
	}
	table = append(table,
		strconv.Itoa(lib.MinInt),
		strconv.Itoa(lib.MaxInt),
		"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
		"123a45",
	)
	for _, arg := range table {
		lib.Function("Atoi", student.Atoi, solutions.Atoi, arg)
	}
}
