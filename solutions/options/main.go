package main

import (
	"fmt"
	"os"
)

func isLatin(r rune) bool {
	for i := 'a'; i <= 'z'; i++ {
		if r == i {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}

	var options [32]bool
	for _, v := range os.Args {
		if len(v) < 2 {
			return
		}
		if v[0] == '-' {
			if v[1] == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			// fill options
			for _, r := range v[1:] {
				if !isLatin(r) {
					fmt.Println("Invalid Option")
					return
				}
				options['z'-r+6] = true
			}
		}
	}
	for i, v := range options {
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
		if v {
			fmt.Print(1)
		} else {
			fmt.Print(0)
		}
	}
	fmt.Println()
}
