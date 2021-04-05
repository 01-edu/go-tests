package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	args := append(rand.MultRandWords(),
		"padinton paqefwtdjetyiytjneytjoeyjnejeyj",
		"ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd",
		"abcdefghij efghijlmnopq",
		"123456 456789",
		"1 1",
		"1 2",
	)

	for i := 0; i < 5; i++ {
		s1 := rand.RandAlnum()
		s2 := rand.RandAlnum() + s1 + rand.RandAlnum()
		args = append(args,
			s1+" "+s2,
			rand.RandAlnum()+" "+rand.RandAlnum(),
		)
	}

	for _, s := range args {
		challenge.Program("inter", strings.Fields(s)...)
	}
}
