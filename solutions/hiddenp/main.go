package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var i int
	if len(os.Args) != 3 {
		return
	}
	for _, r := range os.Args[1] {
		j := strings.Index(os.Args[2][i:], string(r))
		if j == -1 {
			fmt.Println("0")
			return
		}
		i += j + 1
	}
	fmt.Println("1")
}
