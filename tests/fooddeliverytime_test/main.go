package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"burger", "chips", "nuggets", "pizza", "pepper", "  burger", "chi ps", "nuggets  "}
	for _, s := range table {
		challenge.Function("FoodDeliveryTime", student.FoodDeliveryTime, solutions.FoodDeliveryTime, s)
	}
}
