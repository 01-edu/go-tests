package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	a := 0
	b := &a
	n := &b
	student.UltimatePointOne(&n)
	if a != 1 {
		challenge.Fatalf("UltimatePointOne(&n), a == %d instead of 1", a)
	}
}
