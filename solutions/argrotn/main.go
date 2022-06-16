package main

import "fmt"
import "os"
import "strconv"

func rotationByN(s []string, n int) []string {

	str := make([]string, 0)
	if n == 0 {
		return s
	}
	for i := 0; i < len(s); i++ {
		str = append(str, string(s[(i+n)%len(s)]))
	}
	return str
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("")
		return
	}

	args := os.Args[2:]

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Error")
		return
	}
	result := rotationByN(args, n)
	for _, v := range result {
		fmt.Print(v)
		if v != result[len(result)-1] {
			fmt.Print(" ")
		}
	}
}
