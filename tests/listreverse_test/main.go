package main

import (
	student "student"

	"github.com/01-edu/go-tests/solutions"
)

func listPushBack(l *student.List, data interface{}) {
	n := &student.NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func copyList(listStu *student.List) *solutions.List {
	listSol := &solutions.List{}
	it := listStu.Head
	for it != nil {
		solutions.ListPushBack(listSol, it.Data)
		it = it.Next
	}
	return listSol
}

func main() {
	link1 := &solutions.List{}
	link2 := &student.List{}
	table := []solutions.NodeTest{{
		Data: []interface{}{"I", 1, "something", 2},
	}}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		for _, item := range arg.Data {
			listPushBack(link2, item)
			solutions.ListPushBack(link1, item)
		}
		student.ListReverse(link2)
		solutions.ListReverse(link1)

		solutions.ChallengeList("ListReverse", link1, copyList(link2), arg.Data)

		link1 = &solutions.List{}
		link2 = &student.List{}
	}
}
