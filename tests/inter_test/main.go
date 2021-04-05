package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := append(random.MultRandWords(),
		"padinton paqefwtdjetyiytjneytjoeyjnejeyj",
		"ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd",
		"abcdefghij efghijlmnopq",
		"123456 456789",
		"1 1",
		"1 2",
	)

	for i := 0; i < 5; i++ {
		s1 := random.RandAlnum()
		s2 := random.RandAlnum() + s1 + random.RandAlnum()
		args = append(args,
			s1+" "+s2,
			random.RandAlnum()+" "+random.RandAlnum(),
		)
	}

	for _, s := range args {
		challenge.Program("inter", strings.Fields(s)...)
	}
}
