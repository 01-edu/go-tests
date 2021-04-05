package solutions

import "github.com/01-edu/go-tests/lib/is"

func FindNextPrime(nb int) int {
	if is.Prime(nb) {
		return nb
	}
	return FindNextPrime(nb + 1)
}
