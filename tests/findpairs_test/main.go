package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	tests := []struct {
		arr       string
		targetSum string
		want      string
	}{
		{
			arr:       "[1, 2, 3, 4, 5]",
			targetSum: "6",
			want:      "Pairs with sum 6: [[0 4] [1 3]]",
		},
		{
			arr:       "[1, 2, 3, 4, 5, 1]",
			targetSum: "6",
			want:      "Pairs with sum 6: [[0 4] [1 3]]",
		},
		{
			arr:       "[-1, 2, -3, 4, -5]",
			targetSum: "1",
			want:      "Pairs with sum 1: [[0 1] [2 3]]",
		},
		{
			arr:       "[1, 2, 3, 4, 5]",
			targetSum: "-5",
			want:      "Pairs with sum -5: [[0 3] [1 2]]",
		},
		{
			arr:       "[1, 2, 3, 4, 5]",
			targetSum: "10",
			want:      "No pairs found.",
		},
		{
			arr:       "[1, 2, 3, 4, 20, -4, 5]",
			targetSum: "2 5",
			want:      "Invalid target sum.",
		},
		{
			arr:       "[1, 2, 3, 4, 20, -4, 5]",
			targetSum: "l",
			want:      "Invalid target sum.",
		},
		{
			arr:       "[1, 2, 3, 4, 20, p, 5]",
			targetSum: "5",
			want:      "Invalid number: p",
		},
	}

	for _, tc := range tests {
		challenge.Program("findpairs", tc.arr, tc.targetSum)
	}
}
