package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	r := regexp.MustCompile(strings.Join(strings.Split(os.Args[1], ""), ".*"))
	if len(r.FindAllString(os.Args[2], -1)) > 0 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
