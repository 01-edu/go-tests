package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < num; i++ {
		if i%2 == 0 {
			fmt.Println(strconv.Itoa(i) + " ping")
			continue
		}
		fmt.Println(strconv.Itoa(i) + " pong")
	}
}
