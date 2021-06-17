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

func fillOptions(op []string, s string) {
	for _, r := range s {
		if !unicode.Is(unicode.Latin, r) {
			fmt.Println("Invalid Option")
			os.Exit(0)
		}
		op['z'-r+6] = "1"
	}
}

func printOptions(op []string) {
	for i, v := range op {
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		printHelper()
	}

	options := strings.Split(strings.Repeat("0", 32), "")
	for _, v := range os.Args {
		if v[0] == '-' {
			if v[1] == 'h' {
				printHelper()
			}
			fillOptions(options, v[1:])
		}
	}
	printOptions(options)
}
