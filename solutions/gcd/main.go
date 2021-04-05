package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func main() {
	if len(os.Args) == 3 {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[2])
		fmt.Println(gcd(a, b))
	}
}
