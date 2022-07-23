package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	var result byte
	for _, byte := range []byte(os.Args[1]) {
		result += byte
	}
	fmt.Println(result)
}
