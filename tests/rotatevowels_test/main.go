package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	letters := random.Lower + random.Upper + " "
	a := []string{""}
	for i := 0; i < 10; i++ {
		a = append(a, random.RandStr(random.IntBetween(2, 20), letters))
	}

	for _, v := range a {
		challenge.Program("rotatevowels", strings.Fields(v)...)
	}
	challenge.Program("rotatevowels", "Hello World")
	challenge.Program("rotatevowels", "HEllO World", "problem solved")
	challenge.Program("rotatevowels", "str", "shh", "psst")
	challenge.Program("rotatevowels", "happy thoughts", "good luck")
	challenge.Program("rotatevowels", "al's elEphAnt is overly underweight!")
	challenge.Program("rotatevowels", "aEi", "Ou")
}
