package main

import (
	"io/ioutil"
	"os"

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
	table := [][]string{
		{"-c", "23", "src/student/ztail/main.go", file1},
		{"-c", "13", "src/student/ztail/main.go", "fjksdsf", file2},
		{"-c", "5", file1, file2},
		{"-c", "1", "fjksdsf"},
		{"-c", "1", file1, file2},
	}
	for _, args := range table {
		challenge.Program("ztail", args...)
	}
}
