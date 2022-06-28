package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println()
		return
	}
	area, err := strconv.Atoi(args[1])
	if err != nil || area < 0 {
		fmt.Println("Error")
		return
	}
	var result int = int(float32(area*area) * float32(3.14))
	fmt.Println(result)
}
