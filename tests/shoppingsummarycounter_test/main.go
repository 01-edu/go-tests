package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"Burger Burger Water Coffee Water Chips Carrot", "Burger Burger Water Coffe    e Water Chips Carrot", "Burger Burger Water Coffe    e Water Chips Carrot", "32312&%", " ", "", "#$sds "}

	for _, s := range table {
		challenge.Function("ShoppingSummaryCounter", student.ShoppingSummaryCounter, solutions.ShoppingSummaryCounter, s)
	}
}
