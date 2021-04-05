package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	fileName := "quest8.txt"
	if err := ioutil.WriteFile(fileName, []byte("Almost there!!\n"), os.ModePerm); err != nil {
		panic(err)
	}
	table := []string{"", fileName, fileName + " asdsada"}
	for _, s := range table {
		challenge.Program("displayfile", strings.Fields(s)...)
	}
}
