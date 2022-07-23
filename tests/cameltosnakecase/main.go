package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	years := []string{
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCdEf",
		"abcAree",
		"ahe1Abde",
		"ThisIsATest",
	}
	for _, arg := range years {
		challenge.Function("CamelToSnakeCase", student.CamelToSnakeCase, solutions.CamelToSnakeCase, arg)
	}
}
