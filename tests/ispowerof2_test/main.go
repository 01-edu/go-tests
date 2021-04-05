package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][]string{
		{"1"},
		{"2"},
		{"3"},
		{"4"},
		{"1024"},
		{"4096"},
		{"8388608"},
		{"1", "2"},
		{},
	}
	for i := 0; i < 12; i++ {
		args = append(args, []string{strconv.Itoa(random.IntBetween(1, 2048))})
	}
	for _, v := range args {
		challenge.Program("ispowerof2", v...)
	}
}
