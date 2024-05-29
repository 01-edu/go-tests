package main

import (
	"fmt"
	"os"
	student "student"
)

var testCases = []struct {
	in   string
	want string
}{
	{"", "\n"},
	{"             a as", "a\n"},
	{"   f     d", "f\n"},
	{"   asd    ad", "asd\n"},
	{"   salut !!! ", "salut\n"},
	{" salut ! ! !", "salut\n"},
	{"salut ! !", "salut\n"},
}

func main() {
	for _, tc := range testCases {
		got := student.FirstWord(tc.in)
		if got != tc.want {
			fmt.Printf("FirstWord(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}
