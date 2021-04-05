package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	letters := lib.Lower + lib.Upper + " "
	a := []string{""}
	for i := 0; i < 10; i++ {
		a = append(a, lib.RandStr(lib.RandIntBetween(2, 20), letters))
	}

	for _, v := range a {
		lib.Program("rotatevowels", strings.Fields(v)...)
	}
	lib.Program("rotatevowels", "Hello World")
	lib.Program("rotatevowels", "HEllO World", "problem solved")
	lib.Program("rotatevowels", "str", "shh", "psst")
	lib.Program("rotatevowels", "happy thoughts", "good luck")
	lib.Program("rotatevowels", "al's elEphAnt is overly underweight!")
	lib.Program("rotatevowels", "aEi", "Ou")
}
