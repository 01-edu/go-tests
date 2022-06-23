package main

import "fmt"
import "os"

func main() {
	args := os.Args
	if len(args) != 2 || len(args[1]) != 1 {
		return
	}
	ascii := int(args[1][0])
	if ascii >= 65 && ascii <= 90 || ascii >= 97 && ascii <= 122 {
		fmt.Println(ascii)
		return
	}
}
