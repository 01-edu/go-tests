package solutions

import (
	"fmt"
	"strings"
	"unicode"
)

func formatHex(hex string) {
	nbZ := 8 - len(hex)
	if nbZ > 4 {
		hex += strings.Repeat("0", nbZ-4) + " 0000"
	} else {
		hex = hex[:4] + " " + hex[4:] + strings.Repeat("0", nbZ)
	}
	fmt.Print(hex)
}

func PrintMemory(a [10]int) {
	str := ""
	for i, nbr := range a {
		toForm := fmt.Sprintf("%x", nbr)

		if len(toForm) == 1 {
			toForm = "0" + toForm
		}

		formatHex(toForm)
		if ((i+1)%4 == 0 && i != 0) || i == len(a)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}

		if unicode.IsGraphic(rune(nbr)) {
			str += string(rune(nbr))
		} else {
			str += "."
		}
	}
	fmt.Println(str + strings.Repeat(".", 10-len(a)))
}
