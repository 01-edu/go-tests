package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/01-edu/go-tests/lib/challenge"
)

func panicIfNot(target, err error) {
	if err != nil && err != target && !errors.Is(err, target) {
		panic(err)
	}
}

var oldName = "exe"

func test(name string) {
	panicIfNot(nil, exec.Command("go", "build", "-o", oldName, "student/printprogramname").Run())
	panicIfNot(nil, (os.Rename(oldName, name)))
	b, err := exec.Command("./" + name).CombinedOutput()
	panicIfNot(nil, err)
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
