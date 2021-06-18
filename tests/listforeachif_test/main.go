package main

import (
	"reflect"
	"runtime"
	student "student"

	"github.com/01-edu/go-tests/lib/chars"
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

func stringToOneS(node *solutions.NodeL) {
	node.Data = 1
}

func stringToOne(node *student.NodeL) {
	node.Data = 1
}

func main() {
	table := []solutions.NodeTest{{
		Data: []interface{}{"I", 1, "something", 2},
	}, {
		Data: []interface{}{},
	}}

	for i := 0; i < 3; i++ {
		val := solutions.NodeTest{
			Data: solutions.IntToInterface(random.IntSliceBetween(-50, 100)),
		}
		table = append(table, val)
	}

	for i := 0; i < 3; i++ {
		val := solutions.NodeTest{
			Data: solutions.StringToInterface(random.StrSlice(chars.Words)),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}

		for _, item := range arg.Data {
			solutions.ListPushBack(link1, item)
			listPushBack(link2, item)
		}
		solutions.ListForEachIf(link1, stringToOneS, solutions.IsAlNode)
		student.ListForEachIf(link2, stringToOne, student.IsAlNode)

		solutions.ChallengeList("ListForEachIf",
			link1,
			copyList(link2),
			arg.Data,
			runtime.FuncForPC(reflect.ValueOf(stringToOne).Pointer()).Name())
	}
}
