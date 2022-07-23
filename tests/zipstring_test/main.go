package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
)

func getResult(input string) string {
	expectedResult := map[string]string{
		"aaaa":                         "4a",
		"zzzzzZZZZZZ":                  "5z6Z",
		"":                             "",
		"ziiiiipMeee":                  "1z5i1p1M3e",
		"hhellloTthereYouuunggFelllas": "2h1e3l1o1T1t1h1e1r1e1Y1o3u1n2g1F1e3l1a1s",
	}

	return expectedResult[input]
}

func main() {
	args := []string{
		"aaaa",
		"zzzzzZZZZZZ",
		"",
		"ziiiiipMeee",
		"hhellloTthereYouuunggFelllas",
	}
	for _, arg := range args {
		challenge.Function("AlphaCount", student.ZipString, getResult, arg)
	}
}
