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
		var div int
		var mod int
		student.DivMod(a, b, &div, &mod)
		if div != a/b {
			challenge.Fatalf("DivMod(%d, %d, &div, &mod), div == %d instead of %d", a, b, div, a/b)
		}
		if mod != a%b {
			challenge.Fatalf("DivMod(%d, %d, &div, &mod), mod == %d instead of %d", a, b, mod, a%b)
		}
		i++
	}
}
