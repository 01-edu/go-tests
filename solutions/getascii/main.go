package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 && len(os.Args[1]) == 1 && int(os.Args[1][0]) >= 0 && int(os.Args[1][0]) <= 127 {
		fmt.Println(int(os.Args[1][0]))
	}
	fmt.Print("\n")
}
