package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"abc"},
		{"abc", "def"},
		{"abc", "def", "ghi"},
		{"125", "wde"},
		{"125", "wde", "abc"},
	}

	for _, v := range args {
		challenge.Program("argsort", v...)
	}
}
