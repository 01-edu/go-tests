package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(0)
		return
	}
	count := 0
	for _, arg := range os.Args[1:] {
		count += len(arg)
	}
	fmt.Println(count)
}
