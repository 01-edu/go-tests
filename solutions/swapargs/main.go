package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("")
		return
	}
	fmt.Println(os.Args[2] + " " + os.Args[1])
}
