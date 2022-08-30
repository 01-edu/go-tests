package main

import (

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{
		"Loaf of bread",
		"amazing loaf of bread",
		"bread crumbles",
		"",
		"bread slice",
		"bread",
		"This is a loaf of bread",
		"This is a loaf of brea",
		" This is a loaf of bread",
		" This is a loaf of bread ",
		"This is a loaf of bre ",
		"This     is a loaf  of bread",
		"This     is a  loaf of bre ",
		"This is a loaf of bre  ",
		"This is a loaf of br  ",
		"      This     is a l o a f     of   br    ead",
		"                       "}

	for _, s := range table {
		challenge.Function("LoafOfBread", student.LoafOfBread, solutions.LoafOfBread, s)
	}
}
