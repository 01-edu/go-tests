package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

type (
	ListS = solutions.List
	List  = student.List
)

func listPushBack(l *List, data interface{}) {
	n := &student.NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func main() {
	link1 := &ListS{}
	link2 := &List{}
	table := []solutions.NodeTest{{
		Data: []interface{}{3, 2, 1},
	}}

	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		for _, item := range arg.Data {
			listPushBack(link2, item)
			solutions.ListPushBack(link1, item)
		}
		solutionsLast := solutions.ListLast(link1)
		studentLast := student.ListLast(link2)

		if solutionsLast != studentLast {
			challenge.Fatalf("\nlist:%s\n\nListLast() == %v instead of %v\n\n",
				solutions.ListToString(link1.Head), studentLast, solutionsLast)
		}
		link1 = &ListS{}
		link2 = &List{}
	}
}
