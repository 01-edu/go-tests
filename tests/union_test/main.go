package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	s1 := random.Str(chars.Alnum, 13)
	s2 := random.Str(chars.Alnum, 13) + s1 + random.Str(chars.Alnum, 13)

	args := [][]string{
		{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"},
		{"ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"},
		{""},
		{"rien", "cette phrase ne cache rien"},
		{" this is ", " wait shr"},
		{" more ", "then", "two", "arguments"},
		{s1, s2},
		{random.Str(chars.Alnum, 13), random.Str(chars.Alnum, 13)},
	}

	for _, v := range args {
		challenge.Program("union", v...)
	}
}
