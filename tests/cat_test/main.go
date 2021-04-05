package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/01-edu/go-tests/lib"
)

func main() {
	file1 := "quest8.txt"
	file2 := "quest8T.txt"
	if err := ioutil.WriteFile(file1, []byte(lib.RandWords()+"\n"), os.ModePerm); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(file2, []byte(lib.RandWords()+"\n"), os.ModePerm); err != nil {
		panic(err)
	}

	table := []string{file1, file1 + " " + file2, "asd", "", file1 + " abc", "abc " + file2}

	for _, s := range table {
		lib.Program("cat", strings.Fields(s)...)
	}
	lib.ProgramStdin("cat", lib.RandWords()+"\n")
}
