package main

 import (
 	"fmt"
 	"os"
 )

 func main() {
 	if len(os.Args) == 2 && len(os.Args[1]) == 1 {
 		fmt.Println(int(os.Args[1][0]))
 	} else {
 		fmt.Print("\n")
 	}
 }
