package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"hello THERE", "Converted this string into an int", "14/%37", "a b c d e f g h i j k l m n o p q r s t u v w x y z", "a b c d e fghijk l m n o p q r s t u v w x y z ", "A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z", "heLLo There", "  121314", "qwçiodhqwd", " ", ""}

	for _, s := range table {
		challenge.Function("StringToIntSlice", student.StringToIntSlice, solutions.StringToIntSlice, s)
	}
}
