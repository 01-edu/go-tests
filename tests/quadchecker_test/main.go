package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

var testCases = []struct {
	input string
	want  string
}{
	{
		"o",
		"[quadA] [1] [1]\n",
	},
	{
		"/",
		"[quadB] [1] [1]\n",
	},
	{
		"A",
		"[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]\n",
	},
	{
		"A\nA",
		"[quadD] [2] [1]\n",
	},
	{
		"AA",
		"[quadC] [1] [2]\n",
	},
	{
		"A\nC",
		"[quadD] [2] [1] || [quadE] [2] [1]\n",
	},
	{
		"AC",
		"[quadE] [1] [2]\n",
	},
	{
		`o---o
|   |
o---o`,

		"[quadA] [3] [5]\n",
	},
	{
		"0 0",
		"Not a quad function\n",
	},
}

func runProgram(programSource, input string) (string, error) {
	cmd := exec.Command("go", "run", programSource)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func runTests(programSource string) bool {
	for _, tc := range testCases {
		got, err := runProgram(programSource, tc.input)
		if err != nil {
			fmt.Printf("Error running test: %q\n", err)
			return false
		}
		if got != tc.want {
			fmt.Printf("echo %q | quadchecker\n", tc.input)
			fmt.Printf("got: %q\nwant: %q\n", got, tc.want)
			return false
		}
	}
	return true
}

func main() {
	source := path.Join("..", "piscine-go", "quadchecker", "main.go")
	if ok := runTests(source); !ok {
		os.Exit(1)
	}
}
