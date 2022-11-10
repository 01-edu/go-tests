package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][]int{
		{1, 2},
		{0, 0},
		{8423, 9485},
		{-1, -1},
		{2147483647, 3},
		{14342, -1},
	}
	for _, arg := range args {
		challenge.Function("RectPerimeter", student.RectPerimeter, solutions.RectPerimeter, arg[0], arg[1])
	}
}
