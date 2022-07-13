package main

import "fmt"
import "os"
import "strconv"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("")
		return
	}

	args := os.Args[2:]

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Error")
		return
	}

	for i := 0; i < len(args); i++ {
		fmt.Print(string(args[(i+n)%len(args)]))
		if i < len(args)-1 {
			fmt.Print(" ")
		}
	}
}
