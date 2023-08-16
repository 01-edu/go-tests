package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	tests := []string{
		"Hello world!",
		"",
	}
	challenge.HasRandomTests = true
	challenge.FunctionTestSuite("PrintStr", student.PrintStr, PrintStr, tests)
}
