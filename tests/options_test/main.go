package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	var table []string

	table = append(table, "-"+random.RandLower(),
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
		strings.Join([]string{"-", random.RandStr(10, random.Lower)}, ""),
	)

	for _, s := range table {
		challenge.Program("options", strings.Fields(s)...)
	}
}
