package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
)

var tests = []string{
	"hello THERE",
	"14/%37",
	"a b c d e f g h i j k l m n o p q r s t u v w x y z",
	"",
}

func main() {
	challenge.FunctionTestSuite("StringToIntSlice", student.StringToIntSlice, StringToIntSlice, tests)
}
