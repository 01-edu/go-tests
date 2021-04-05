package main

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
)

func main() {
	for i := 0; i < 25; i++ {
		challenge.Program("costumeprofit",
			strconv.Itoa(rand.IntBetween(0, 1000)),
			strconv.Itoa(rand.IntBetween(0, 1000)),
			strconv.Itoa(rand.IntBetween(0, 1000)),
			strconv.Itoa(rand.IntBetween(0, 1000)),
			strconv.Itoa(rand.IntBetween(0, 1000)),
			strconv.Itoa(rand.IntBetween(0, 1000)),
		)
	}
}
