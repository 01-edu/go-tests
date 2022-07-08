package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{"Abc", "Bcd"},
		{"Hello", " World!"},
		{"talented ", "student"},
		{"A", "1"},
		{"Hello", "World"},
		{""},
		{"he", "llo"},
		{"1", "93"},
		{"he", "llo", " world!"},
		{" ", "		"},
		{"1 + ", "9 = 10"},
		{"", ""},
		{"01", "zone", "Oujda", "Morroco"},
		{},
		{" abc", "def12", "ghk34"},
	}
	for _, arg := range args {
		challenge.Program("concat2args", arg...)
	}
}
