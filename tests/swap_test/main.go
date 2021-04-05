package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	i := 0
	for i < 30 {
		a := random.Int()
		b := random.Int()
		aCopy := a
		bCopy := b
		student.Swap(&a, &b)
		if a != bCopy {
			challenge.Fatalf("Swap(%d, %d), a == %d instead of %d", aCopy, bCopy, a, bCopy)
		}
		if b != aCopy {
			challenge.Fatalf("Swap(%d, %d), b == %d instead of %d", aCopy, bCopy, b, aCopy)
		}
		i++
	}
}
