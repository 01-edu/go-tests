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
		var div int
		var mod int
		student.DivMod(a, b, &div, &mod)
		if div != a/b {
			challenge.Fatalf("DivMod(%d, %d, &div, &mod), div == %d instead of %d", a, b, div, a/b)
		}
		if mod != a%b {
			challenge.Fatalf("DivMod(%d, %d, &div, &mod), mod == %d instead of %d", a, b, mod, a%b)
		}
	}
}
