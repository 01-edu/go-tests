package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/go-tests/solutions"
)

func main() {
	root := &solutions.TreeNode{Data: "08"}
	rootS := &student.TreeNode{Data: "08"}
	pos := []string{
		"x",
		"z",
		"y",
		"t",
		"r",
		"q",
		"01",
		"b",
		"c",
		"a",
		"d",
	}

	for _, arg := range pos {
		root = solutions.BTreeInsertData(root, arg)
		rootS = student.BTreeInsertData(rootS, arg)
	}

	solutions.ChallengeTree("BTreeApplyInorder", solutions.BTreeApplyInorder, student.BTreeApplyInorder, root, rootS, fmt.Println)
}
