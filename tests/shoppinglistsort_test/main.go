package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	table := []string{"Banana", "Mushroom", "Salt", "Pepper","Tea", "Milk", " Milk", " ", "Peanut-Butter", "Banana  "}

	for _, s := range table {
		challenge.Function("ShoppingListSort", student.ShoppingListSort, solutions.ShoppingListSort, s)
	}
}
