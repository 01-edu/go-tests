package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := []string{
		"CDCCDDCDCD",
		"CDDDDCCCDC",
		"DDDDCCCC",
		"CDCCCDDCDD",
		"CDCDCDCDCDCDCDCD",
		"CCCDDDCDCCCCDC",
		"DDDDDDDDDDDDDDDDDDDDDDDDCCCCCCCCCCCCCCCCCCCCCCCC",
		"DDDDCDDDDCDDDCCC",
	}

	for i := 0; i < 15; i++ {
		s := ""
		chunks := random.IntBetween(5, 10)
		for j := 0; j < chunks; j++ {
			countC := random.IntBetween(1, 5)
			countD := random.IntBetween(1, 5)
			tmpC := countC
			tmpD := countD
			for tmpC > 0 || tmpD > 0 {
				letter := random.Str("CD", 1)
				if tmpC > 0 && letter == "C" {
					tmpC--
					s += letter
				} else if tmpD > 0 && letter == "D" {
					tmpD--
					s += letter
				}
			}

			tmpC = countC
			tmpD = countD
			for tmpC > 0 || tmpD > 0 {
				letter := random.Str("CD", 1)
				if tmpC > 0 && letter == "D" {
					tmpC--
					s += letter
				} else if tmpD > 0 && letter == "C" {
					tmpD--
					s += letter
				}
			}
		}
		args = append(args, s)
	}

	for _, arg := range args {
		challenge.Program("balancedstring", arg)
	}
}
