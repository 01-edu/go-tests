package main

import (
	"fmt"
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"Hello World",
		"HelloWorldB",
		"ClipClapWalk",
		"ClipClapWalkClipClapWalk",
		"#@13word",
		"Wron2Word",
		"Oneword",
		"",
		" ",
		"RightWordLargeWorldALotOfWords",
	}
	for _, arg := range args {
		challenge.Function("SetSpace", student.SetSpace, solutions.SetSpace, arg)
	}
}
