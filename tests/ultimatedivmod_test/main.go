package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	i := 0
	for i < lib.SliceLen {
		a := lib.RandInt()
		b := lib.RandInt()
		aCopy := a
		bCopy := b
		div := a / b
		mod := a % b
		student.UltimateDivMod(&a, &b)
		if a != div {
			challenge.Fatalf("DivMod(%d, %d), a == %d instead of %d", aCopy, bCopy, a, div)
		}
		if b != mod {
			challenge.Fatalf("DivMod(%d, %d), b == %d instead of %d", aCopy, bCopy, b, mod)
		}
		i++
	}
}
