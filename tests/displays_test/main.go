package main

import "github.com/01-edu/go-tests/lib/challenge"

func main() {
	args := []string{
		"dsfdz",
		"",
		"1",
		"2",
	}
	for _, s := range args {
		challenge.Program("displays", s)
	}
}
