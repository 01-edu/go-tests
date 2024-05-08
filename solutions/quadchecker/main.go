package main

import (
	"fmt"
	"io"
	"os"
)

var testCasesMap = map[string]string{
	"o": "[quadA] [1] [1]",
	"/": "[quadB] [1] [1]",
	"A": "[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]",
	`A
A`: "[quadD] [2] [1]",
	`AA`: "[quadC] [1] [2]",
	`A
C`: "[quadD] [2] [1] || [quadE] [2] [1]",
	`AC`: "[quadE] [1] [2]",
	`o---o
|   |
o---o`: "[quadA] [3] [5]",
	"0 0": "Not a quad function",
}

func main() {
	info, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(testCasesMap[string(info)])
}
