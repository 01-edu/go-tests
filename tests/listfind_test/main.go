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
		Data: []interface{}{"hello", "hello1", "hello2", "hello3"},
	}}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}
		for _, item := range arg.Data {
			solutions.ListPushBack(link1, item)
			listPushBack(link2, item)
		}

		if len(arg.Data) > 0 {
			stuValue := student.ListFind(link2, arg.Data[(len(arg.Data)-1)/2], student.CompStr)
			solValue := solutions.ListFind(link1, arg.Data[(len(arg.Data)-1)/2], solutions.CompStr)
			if stuValue == nil && solValue != nil {
				challenge.Fatalf("ListFind(ref: %s) == nil instead of %s\n",
					arg.Data[(len(arg.Data)-1)/2],
					*solValue)
			}
			if stuValue != nil || solValue != nil {
				if *stuValue != *solValue {
					challenge.Fatalf("ListFind(ref: %s) == %v instead of %s\n",
						arg.Data[(len(arg.Data)-1)/2],
						*stuValue,
						*solValue)
				}
			}
		}
	}
}
