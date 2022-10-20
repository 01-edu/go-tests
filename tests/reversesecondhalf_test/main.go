package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"hello wold", "", "Kimetsu no Yaiba", "Z", "email123@live.fr", "write  ==> 45m$", "-552", "w58tw7474abc", "fifa world cup `2022`"}
	for _, s := range table {
		challenge.Function("ReverseSecondHalf", student.ReverseSecondHalf, solutions.ReverseSecondHalf, s)
	}
}
