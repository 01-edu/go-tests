package main

import (
	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"hello", "you"},
		{"   only  it's harder   "},
		{"you   see   it's   easy   to   display    the     same  thing"},
	}

	args = append(args, lib.MultRandWords())

	for _, v := range args {
		challenge.Program("expandstr", v...)
	}
	challenge.Program("expandstr")
}
