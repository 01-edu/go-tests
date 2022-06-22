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
	for i := len(os.Args) - 1; i >= 1; i-- {
		fmt.Print(os.Args[i])
		if i != 1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
