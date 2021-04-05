package solutions

import (
	"fmt"

	"github.com/01-edu/go-tests/lib/base"
)

func PrintNbrBase(n int, b string) {
	if base.IsValid(b) {
		length := len(b)
		sign := 1
		rbase := []rune(b)
		if n < 0 {
			fmt.Print("-")
			sign = -1
		}
		if n < length && n >= 0 {
			fmt.Printf("%c", rbase[n])
		} else {
			PrintNbrBase(sign*(n/length), b)
			fmt.Printf("%c", rbase[sign*(n%length)])
		}
	} else {
		fmt.Print("NV")
	}
}
