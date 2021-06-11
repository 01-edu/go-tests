package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func copyList(listStu *student.List, listSol *solutions.List) *solutions.List {
	for it := listStu.Head; it != nil; it = it.Next {
		solutions.ListPushBack(listSol, it.Data)
	}
	return listSol
}

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
	link1 := &solutions.List{}
	link2 := &student.List{}

	table := []solutions.NodeTest{{
		Data: []interface{}{"I", 1, "something", 2},
	}}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		for i := 0; i < len(arg.Data); i++ {
			listPushBack(link2, arg.Data[i])
			solutions.ListPushBack(link1, arg.Data[i])
		}
		solutions.ListClear(link1)
		student.ListClear(link2)

		if link2.Head != nil {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListClear() == %v instead of %v\n\n",
				solutions.ListToString(copyList(link2, &solutions.List{}).Head),
				solutions.ListToString(link1.Head), link2.Head, link1.Head)
		}
	}
}
