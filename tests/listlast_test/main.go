package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
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

func main() {
	table := []solutions.NodeTest{{
		Data: []interface{}{3, 2, 1},
	}}

	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}
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
	}
}
