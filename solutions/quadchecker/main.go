package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	info, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	x, y := 0, 0

	for _, r := range info {
		if r == 10 {
			y++
		}
	}
	for i := 0; info[i] != 10; i++ {
		x++
	}
	if info[0] == 'o' {
		fmt.Printf("[quadA] [%d] [%d]\n", x, y)
	}
	if info[0] == '/' {
		fmt.Printf("[quadB] [%d] [%d]\n", x, y)
	}
	if x == 1 && y == 1 && info[0] == 'A' {
		fmt.Printf("[quadC] [%d] [%d] || ", x, y)
		fmt.Printf("[quadD] [%d] [%d] || ", x, y)
		fmt.Printf("[quadE] [%d] [%d]\n", x, y)
	} else if x == 1 && info[0] == 'A' && info[2*(y-x)] == 'A' {
		fmt.Printf("[quadD] [%d] [%d]\n", x, y)
	} else if x == 1 && info[0] == 'A' && info[2*(y-x)] == 'C' {
		fmt.Printf("[quadC] [%d] [%d] || ", x, y)
		fmt.Printf("[quadE] [%d] [%d]\n", x, y)
	} else if y == 1 && info[0] == 'A' && info[x-1] == 'C' {
		fmt.Printf("[quadD] [%d] [%d] || ", x, y)
		fmt.Printf("[quadE] [%d] [%d]\n", x, y)
	} else if y == 1 && info[0] == 'A' && info[x-1] == 'A' {
		fmt.Printf("[quadC] [%d] [%d]\n", x, y)
	} else if info[0] == 'A' && info[x-1] == 'A' && info[x*y-1] == 'C' && info[x*y+x-2] == 'C' {
		fmt.Printf("[quadC] [%d] [%d]\n", x, y)
	} else if info[0] == 'A' && info[x-1] == 'C' && info[x*y-1] == 'A' && info[x*y+x-2] == 'C' {
		fmt.Printf("[quadD] [%d] [%d]\n", x, y)
	} else if info[0] == 'A' && info[x-1] == 'C' && info[x*y-1] == 'C' && info[x*y+x-2] == 'A' {
		fmt.Printf("[quadE] awdawd[%d] [%d]\n", x, y)
	}
}
