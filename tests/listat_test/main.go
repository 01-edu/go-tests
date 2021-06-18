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
	type nodeTest struct {
		data []interface{}
		pos  int
	}

	table := []nodeTest{{
		data: []interface{}{"I", 1, "something", 2},
	}, {
		data: []interface{}{"I", 1, "something", 2},
		pos:  random.IntBetween(1, 4),
	}}

	for i := 0; i < 4; i++ {
		table = append(table, nodeTest{
			data: solutions.IntToInterface(random.IntSliceBetween(-50, 100)),
			pos:  random.IntBetween(1, 12),
		})
	}

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}
		for _, item := range arg.data {
			listPushBack(link2, item)
			solutions.ListPushBack(link1, item)
		}

		solutions.ChallengeList("ListAt",
			&solutions.List{Head: solutions.ListAt(link1.Head, arg.pos)},
			copyList(&student.List{Head: student.ListAt(link2.Head, arg.pos)}),
			arg.data, arg.pos)
	}
}
