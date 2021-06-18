package main

import (
	"reflect"
	"runtime"
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
	for it := listStu.Head; it != nil; it = it.Next {
		solutions.ListPushBack(listSol, it.Data)
	}
	return listSol
}

func main() {
	table := []solutions.NodeTest{{
		Data: []interface{}{"I", 1, "something", 2},
	}}
	table = solutions.ElementsToTest(table)

	for _, arg := range table {
		link1 := &solutions.List{}
		link2 := &student.List{}

		for _, item := range arg.Data {
			solutions.ListPushBack(link1, item)
			listPushBack(link2, item)
		}
		student.ListForEach(link2, student.Add2_node)
		solutions.ListForEach(link1, solutions.Add2_node)

		solutions.ChallengeList("ListForEach",
			link1,
			copyList(link2),
			arg.Data,
			runtime.FuncForPC(reflect.ValueOf(student.Add2_node).Pointer()).Name())
	}
}
