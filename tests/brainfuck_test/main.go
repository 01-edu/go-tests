package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	// individual tests 1)Hello World! 2)Hi 3)abc 4)ABC

	args := []string{
		"++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.",
		"+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>.",
		"++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++.",
		"ld++++++++++++++++++++++++++++++++++++++++++++this+is++a++++comment++++++++++++++[>d+<-]>.+.+.>++++++++++.",
		strings.Join(
			[]string{
				"ld++++++++++++++++++++++++++++++++++++++++++++this+is++a++++comment++++++++++++++[>d+<-]>.+",
				random.Str(".+", random.IntBetween(1, 10)),
				".+.>++++++++++.",
			},
			"",
		),
	}
	for _, v := range args {
		challenge.Program("brainfuck", v)
	}
}
