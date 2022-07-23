package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	args := [][]string{
		{},
		{""},
		{" "},
		{"		"},
		{"One", "ring!"},
		{"Pirate ipsum"},
		{
			"It is when pirates count their booty that they become mere thieves. ",
			"Shiver me timbers. ",
			"Arrrrrrrr Ahoy!",
			"lets trouble the water!",
		},
		{"z"},
		{"testing spaces and #!*"},
		{"more", "than", "three", "arguments"},
		{"What", "are", "you", "doing?"},
		{"1"},
		{"née", "changèd", "entrepôt", "Zoë", "Māori", "drŏll", "über", "Señor", "João"},
	}
	for _, v := range args {
		challenge.Program("canyoucount", v...)
	}
}
