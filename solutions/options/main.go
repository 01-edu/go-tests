package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func printHelper() {
	fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
	os.Exit(0)
}

func printOptions(op []string) {

}

func main() {
	if len(os.Args) < 2 {
		printHelper()
	}

	options := strings.Split("00000000000000000000000000000000", "")
	for _, v := range os.Args {
		if len(v) < 2 {
			return
		}
		if v[0] == '-' {
			if v[1] == 'h' {
				printHelper()
			}
			// fill options
			for _, r := range v[1:] {
				if !unicode.Is(unicode.Latin, r) {
					fmt.Println("Invalid Option")
					os.Exit(0)
				}
				options['z'-r+6] = "1"
			}
		}
	}
	for i, v := range options {
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
