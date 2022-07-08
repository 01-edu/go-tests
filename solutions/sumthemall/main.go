package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(0)
		return
	}
	result := 0
	for _, arg := range os.Args[1:] {
		num, err := strconv.Atoi(arg)
		if num >= MaxInt || num <= MinInt || err != nil {
			fmt.Println(0)
			return
		}
		result += num
	}
	fmt.Println(result)
}
