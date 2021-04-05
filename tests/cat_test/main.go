package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	file1 := "quest8.txt"
	file2 := "quest8T.txt"
	if err := ioutil.WriteFile(file1, []byte(random.Str(chars.Words, 13)+"\n"), os.ModePerm); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(file2, []byte(random.Str(chars.Words, 13)+"\n"), os.ModePerm); err != nil {
		panic(err)
	}
	args := []string{
		file1,
		file1 + " " + file2,
		"asd",
		"",
		file1 + " abc",
		"abc " + file2,
	}
	for _, s := range args {
		challenge.Program("cat", strings.Fields(s)...)
	}
	challenge.ProgramStdin("cat", random.Str(chars.Words, 13)+"\n")
}
