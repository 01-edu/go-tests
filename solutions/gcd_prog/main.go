package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	if len(os.Args) == 3 {
		a, _ := strconv.Atoi(os.Args[1])
		b, _ := strconv.Atoi(os.Args[2])
		fmt.Println(lib.GCD(a, b))
	}
}
