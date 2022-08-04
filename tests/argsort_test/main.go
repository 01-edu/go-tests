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
		{"a b c"},
		{"   abc"},
		{"    hello"},
		{"hello    "},
		{"sto  ry"},
		{"  @\\!\\n"},
		{"9^_\\`abc"},
		{"	\n"},
		{"\t\v\n"},
		{"%Hey|}"},
		{"123	$_$%d"},
	}
	for _, v := range args {
		challenge.Program("argsort", v...)
	}
}
