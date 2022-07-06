package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println()
		return
	}
	for i := 1; i < len(os.Args); i++ {
		if i != len(os.Args)-1 {
			fmt.Print(os.Args[i+1])
			fmt.Print(" ")
		} else {
			fmt.Print(os.Args[1])
		}
	}
	fmt.Println()
}
