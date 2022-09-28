package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	slice := [][]string{
		{"desserts", "mains", "drinks", "starters"},
		{"des  rts", "mains ", "  drinks", "starters  "},
		{"", "desrts", "mins", "drins", "strters"},
		{" ", "des serts", "ma ins", "drink s", "st arters"},
		{"desserts", "mai  ns", "dr   inks", "st3arters"},
		{"desserts", "mains", "drinks", "starters"},
		{""},
		{"+55", "*@Â£%D", "edr8927", "		| cat -e"},
		{"741", "852", "4", "58", "87", "03", "-96"},
		{"kd", "jq", "hel$lo", "a", "B", "9W"},
	}

	for _, s := range slice {
		challenge.Function("ReverseMenuIndex", student.ReverseMenuIndex, solutions.ReverseMenuIndex, s)
	}
}
