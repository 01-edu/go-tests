package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(random.Int())
	}
	table = append(table,
		strconv.Itoa(random.MinInt),
		strconv.Itoa(random.MaxInt),
		"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"+657",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
		"123a45",
		"+1234",
		"-1234",
		"+123456",
		"-123456",
	)
	for _, arg := range table {
		challenge.Function("Atoi", student.Atoi, solutions.Atoi, arg)
	}
}
