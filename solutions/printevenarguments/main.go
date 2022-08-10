package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 2 {
		fmt.Println()
		return
	}
	for i, arg := range os.Args[1:] {
		if i%2 != 0 {
			fmt.Println(arg)
		}
	}
}
