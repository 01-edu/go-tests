package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"-" + random.Str(chars.Lower, 13),
		"-",
		" ",
		"-%",
		"-?",
		"-=",
		"-abc -ijk",
		"-something-",
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
		"-" + random.Str(chars.Lower, 13),
	}
	for _, s := range args {
		challenge.Program("options", strings.Fields(s)...)
	}
}
