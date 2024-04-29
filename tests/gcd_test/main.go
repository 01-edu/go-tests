package main

import (
	"fmt"
	"os"
	student "student"
)

func main() {
	testCases := []struct {
		a    uint
		b    uint
		want uint
	}{
		{42, 10, 2},
		{42, 12, 6},
		{14, 77, 7},
		{17, 3, 1},
		{12, 23, 1},
		{25, 15, 5},
		{23043, 122, 1},
		{11, 77, 11},
	}

	for _, tc := range testCases {
		got := student.Gcd(tc.a, tc.b)
		if got != tc.want {
			fmt.Printf("Gcd(%d, %d) = %d instead of %d\n", tc.a, tc.b, got, tc.want)
			os.Exit(1)
		}
	}
}
