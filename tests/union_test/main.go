package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	s1 := random.RandAlnum()
	s2 := strings.Join([]string{random.RandAlnum(), s1, random.RandAlnum()}, "")

	args := [][]string{
		{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"},
		{"ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"},
		{""},
		{"rien", "cette phrase ne cache rien"},
		{" this is ", " wait shr"},
		{" more ", "then", "two", "arguments"},
		{s1, s2},
		{random.RandAlnum(), random.RandAlnum()},
	}

	for _, v := range args {
		challenge.Program("union", v...)
	}
}
