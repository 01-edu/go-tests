package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib"
	"github.com/01-edu/go-tests/solutions"
)

func randomSize() []int {
	var randSlice []int
	for i := 0; i <= lib.RandIntBetween(0, 20); i++ {
		randSlice = append(randSlice, lib.RandInt())
	}
	return randSlice
}

func main() {
	type node struct {
		slice []int
		ch    int
	}
	table := []node{}

	for i := 0; i <= 7; i++ {
		value := node{
			slice: randomSize(),
			ch:    lib.RandIntBetween(0, 10),
		}
		table = append(table, value)
	}
	table = append(table, node{
		slice: []int{},
		ch:    0,
	}, node{
		slice: []int{1, 2, 3, 4, 5, 6, 7, 8},
		ch:    0,
	})
	for _, args := range table {
		lib.Function("Chunk", student.Chunk, solutions.Chunk, args.slice, args.ch)
	}
}
