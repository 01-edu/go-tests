package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

var tests = []string{
	"0",
	"12345",
	"0000012345",
	"01",
	"49819",
	"2147483647",
	"a89516",
	"!01",
	"748q3647",
}

var results = []int{
	0,
	12345,
	12345,
	1,
	49819,
	2147483647,
	0,
	0,
	0,
}

func main() {
	challenge.HasRandomTests = true
	random.StringCharset = "0123456789"
	random.MaxStringLen = 9
	challenge.FunctionTestSuite("BasicAtoi", student.BasicAtoi, BasicAtoi, tests)
}
