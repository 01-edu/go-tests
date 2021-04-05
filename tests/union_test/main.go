package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	s1 := rand.RandAlnum()
	s2 := strings.Join([]string{rand.RandAlnum(), s1, rand.RandAlnum()}, "")

	args := [][]string{
		{"zpadinton", "paqefwtdjetyiytjneytjoeyjnejeyj"},
		{"ddf6vewg64f", "gtwthgdwthdwfteewhrtag6h4ffdhsd"},
		{""},
		{"rien", "cette phrase ne cache rien"},
		{" this is ", " wait shr"},
		{" more ", "then", "two", "arguments"},
		{s1, s2},
		{rand.RandAlnum(), rand.RandAlnum()},
	}

	for _, v := range args {
		challenge.Program("union", v...)
	}
}
