package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	slice := [][]string{
		{"Banana", "Mushroom", "Salt", "Pepper", "Tea", "Milk"},
		{"Bna", "Mushroom", "Sabdfdflt", "Peper", "Tea", "Milvsdgvdk"},
		{"", "Musbfgbfgom", "St", "Pepper", "Tea", "Miluiyk"},
		{"", "Mushroom", "Salt", "Pepper", "a", "Milk"},
		{"Banuana", "Mushroom", "185+", "Pepp8er", "", "Milkhfghrhr"},
		{"Bna", "Mm", "Sa", "Peprhtrper", "Tea", "Milk"},
		{""},
		{"+55", "*@Â£%D", "edr8927", "		| cat -e"},
		{"741", "852", "4", "58", "87", "03", "-96"},
		{"kd", "jq", "hel$lo", "a", "B", "9W"},
	}

	for _, s := range slice {
		challenge.Function("ShoppingListSort", student.ShoppingListSort, solutions.ShoppingListSort, s)
	}
}
