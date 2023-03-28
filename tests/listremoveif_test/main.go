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
	for it := listStu.Head; it != nil; it = it.Next {
		solutions.ListPushBack(listSol, it.Data)
	}
	return listSol
}

func main() {
	var index int

	table := []solutions.NodeTest{
		{Data: []interface{}{"hello", "hello1", "hello2", "hello3"}},
		{Data: []interface{}{1, "Hello", 1, "There", 1, 1, "How", 1, "are", "you", 1}},
		{Data: []interface{}{0, "is", 1, "there", 1, 1, "some", 1, "coins", "here", 1}},
	}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}

		for _, item := range arg.Data {
			solutions.ListPushBack(link1, item)
			listPushBack(link2, item)
		}

		index = random.IntBetween(0, len(arg.Data)-1)

		if link1.Head != nil && link2.Head != nil {
			chosenOne := arg.Data[index]
			if link1.Head.Data == 1 {
				chosenOne = 1
			}
			student.ListRemoveIf(link2, chosenOne)
			solutions.ListRemoveIf(link1, chosenOne)
			solutions.ChallengeList("ListRemoveIf", link1, copyList(link2), arg.Data, chosenOne)
		}
	}
	// Remove more than one node at the begging
	link1 := &solutions.List{}
	link2 := &student.List{}

	for _, item := range table[2].Data {
		solutions.ListPushBack(link1, item)
		listPushBack(link2, item)
	}
	student.ListRemoveIf(link2, 0)
	student.ListRemoveIf(link2, "is")
	solutions.ListRemoveIf(link1, 0)
	solutions.ListRemoveIf(link1, "is")
	solutions.ChallengeList("ListRemoveIf", link1, copyList(link2), table[2].Data, 0, "is")
}
