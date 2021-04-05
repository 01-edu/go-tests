package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			panic(err)
		}
	} else {
		for _, arg := range os.Args[1:] {
			b, err := ioutil.ReadFile(arg)
			if err != nil {
				challenge.Fatalln("ERROR:", err)
			}
			os.Stdout.Write(b)
		}
	}
}
