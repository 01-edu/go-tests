package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	var table []string

	table = append(table, "-"+lib.RandLower(),
		" ",
		"-%",
		"-?",
		"-=",
		"-abc -ijk",
		"-z",
		"-abc -hijk",
		"-abcdefghijklmnopqrstuvwxyz",
		"-abcdefgijklmnopqrstuvwxyz",
		"-eeeeee",
		"-hhhhhh",
		"-h",
		"-hz",
		"-zh",
		"-z -h",
		strings.Join([]string{"-", lib.RandStr(10, lib.RuneRange('a', 'z'))}, ""),
	)

	for _, s := range table {
		challenge.Program("options", strings.Fields(s)...)
	}
}
