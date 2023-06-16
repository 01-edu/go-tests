package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"1234556789", "QKplq%QSw", "", "Kimetsu no Yaiba", "Z", "email123@live.fr", "8595485-52", "-552", "abc", "w58tw7474abc", "fifa world cup `2022`"}

	for _, s := range table {
		challenge.Function("ThirdTimeIsACharm", student.ThirdTimeIsACharm, solutions.ThirdTimeIsACharm, s)
	}
}
