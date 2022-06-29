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
	result := ""
	for _, arg := range os.Args[1:] {
		result += arg
	}
	fmt.Println(result)
}
