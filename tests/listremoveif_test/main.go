package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/random"
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
	var index int
	table := []solutions.NodeTest{{
		Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
	}}

	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		for _, item := range arg.Data {
			solutions.ListPushBack(link1, item)
			listPushBack(link2, item)
		}

		index = random.IntBetween(0, len(arg.Data)-1)
		if link1.Head != nil && link2.Head != nil {
			chosenOne := arg.Data[index]
			student.ListRemoveIf(link2, chosenOne)
			solutions.ListRemoveIf(link1, chosenOne)
			solutions.ChallengeList("ListRemoveIf", link1, copyList(link2), arg.Data, chosenOne)
		} else {
			student.ListRemoveIf(link2, 1)
			solutions.ListRemoveIf(link1, 1)
			solutions.ChallengeList("ListRemoveIf", link1, copyList(link2), arg.Data, 1)
		}

		link1 = &solutions.List{}
		link2 = &student.List{}
	}
}
