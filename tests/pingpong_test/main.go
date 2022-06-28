package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{""},
		{"0"},
		{"3", "5"},
		{"5"},
		{"9"},
		{"24"},
		{"99"},
		{"100"},
	}
	for i := 0; i < 10; i++ {
		rand := []string{strconv.Itoa(random.IntBetween(0, 99))}
		args = append(args, rand)
	}
	for _, arg := range args {
		challenge.Program("pingpong", arg...)
	}
}
