package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"abc"},
		{"abc", "def"},
		{"125"},
		{"125", "wde"},
		{"def"},
		{""},
		{"123qa"},
		{"	 HELLO!"},
		{" kowz"},
	}
	for _, v := range args {
		challenge.Program("argsort", v...)
	}
}
