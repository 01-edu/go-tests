package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
)

func removeBinary(s ...string) {
	for _, c := range s {
		e := os.Remove(c)
		if e != nil {
			fmt.Printf("error: %q\n", e)
		}
	}
}

func buildBin(path, binName string) error {
	cmd := exec.Command("go", "build", "-o", binName, path)
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil

}

func runBin(binName, input string) (string, error) {
	cmd := exec.Command("./" + binName)
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

// only one output no matter the arguments
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

func runTests(binName string) bool {
	for _, tc := range testCases {
		got, err := runBin(binName, tc.input)
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
	binName := "quadchecker"
	if err := buildBin(path.Join("..", "piscine-go", "quadchecker", "main.go"), binName); err != nil {
		fmt.Printf("Error building binary: %q\n", err)
		os.Exit(1)
	}
	defer removeBinary(binName)

	if ok := runTests(binName); !ok {
		os.Exit(1)
	}
}
