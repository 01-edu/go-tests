package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func add0(i int) {
	fmt.Println(i)
}

func add1(i int) {
	fmt.Println(i + 1)
}

func add2(i int) {
	fmt.Println(i + 2)
}

func add3(i int) {
	fmt.Println(i + 3)
}

func main() {
	functions := []func(int){add0, add1, add2, add3}

	type node struct {
		f func(int)
		a []int
	}

	table := []node{{
		f: add0,
		a: []int{1, 2, 3, 4, 5, 6},
	}}

	// 15 random slice of random ints with a random function from selection
	for i := 0; i < 15; i++ {
		function := functions[random.IntBetween(0, len(functions)-1)]
		table = append(table, node{
			f: function,
			a: random.IntSliceBetween(-1000000, 1000000),
		})
	}

	for _, arg := range table {
		challenge.Function("ForEach", student.ForEach, solutions.ForEach, arg.f, arg.a)
	}
}
