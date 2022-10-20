package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	counter := 0
	if length > 2 {
	} else {
		lenString := len(os.Args[1])
		for i := 0; i < lenString; i++ {
			if os.Args[1][i] != 97 {
				if counter == lenString-1 {
					fmt.Print("!(Contains the letter)\n")
					break
				}
				counter++

			} else if os.Args[1][i] == 97 {
				fmt.Print("Contains the letter\n")
				break
			}
		}
	}
}
