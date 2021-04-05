package main

import (
	"os"
	"os/exec"

	"github.com/01-edu/go-tests/lib/challenge"
)

var oldName = "exe"

func test(name string) {
	if err := os.Rename(oldName, name); err != nil {
		challenge.Fatalln(err)
	}
	b, err := exec.Command("./" + name).CombinedOutput()
	if err != nil {
		challenge.Fatalln(err)
	}
	if string(b) != name+"\n" {
		challenge.Fatalln("Failed to print the program name :", string(b))
	}
	oldName = name
}

func main() {
	test("choumi")
	test("🤦🏻‍♀️")
	test("€")
}
