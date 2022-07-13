package main

import "fmt"
import "strconv"
import "os"

func main() {
	if len(os.Args) < 3 {
		fmt.Println()
		return
	}
	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		return
	}
	var max int = i
	var min int = i
	for _, arg := range os.Args[1:] {
		i, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Error")
			return
		}
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	fmt.Println(min, max)
}
