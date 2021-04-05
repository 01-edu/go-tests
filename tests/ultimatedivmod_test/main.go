package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	for i := 0; i < 8; i++ {
		a := random.Int()
		b := random.Int()
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
	}
}
