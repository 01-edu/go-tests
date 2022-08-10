package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println()
		return
	}
	args := os.Args[1:]

	if len(args)%2 == 0 {
		fmt.Println(args[len(args)/2-1] + " " + args[len(args)/2])
	} else {
		fmt.Println(args[len(args)/2])
	}
}
