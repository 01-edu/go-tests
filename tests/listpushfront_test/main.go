package main

import (
	student "student"

	"github.com/01-edu/go-tests/solutions"
)

func copyList(listStu *student.List) *solutions.List {
	listSol := &solutions.List{}
	for it := listStu.Head; it != nil; it = it.Next {
		solutions.ListPushBack(listSol, it.Data)
	}
	return listSol
}

func main() {
	table := []solutions.NodeTest{{
		Data: []interface{}{"Hello", "man", "how are you"},
	}}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}

		for _, item := range arg.Data {
			student.ListPushFront(link2, item)
			solutions.ListPushFront(link1, item)
		}
		solutions.ChallengeList("ListPushFront", link1, copyList(link2), arg.Data)
	}
}
