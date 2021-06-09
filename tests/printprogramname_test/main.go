package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/01-edu/go-tests/lib/challenge"
)

func expect(target, err error) {
	if err != nil && err != target && !errors.Is(err, target) {
		panic(err)
	}
}

var oldName = "exe"

func test(name string) {
	expect(nil, exec.Command("go", "build", "-o", oldName, "student/printprogramname").Run())
	expect(nil, (os.Rename(oldName, name)))
	b, err := exec.Command("./" + name).CombinedOutput()
	expect(nil, err)
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
