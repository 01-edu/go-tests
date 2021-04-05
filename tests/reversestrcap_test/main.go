package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{"First SMALL TesT"},
		{"SEconD Test IS a LItTLE EasIEr", "bEwaRe IT'S NoT HARd WhEN ", " Go a dernier 0123456789 for the road e"},
		{""},
	}

	for i := 0; i < 15; i++ {
		args = append(args, random.StrSlice(chars.Alnum))
	}

	for _, v := range args {
		challenge.Program("reversestrcap", v...)
	}
	challenge.Program("reversestrcap")
}
