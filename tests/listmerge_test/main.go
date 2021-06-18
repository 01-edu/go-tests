package main

import (
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

func main() {
	type nodeTest struct {
		data1 []interface{}
		data2 []interface{}
	}
	table := []nodeTest{{
		data1: []interface{}{},
		data2: []interface{}{"a", 1},
	}}
	// empty list
	table = append(table,
		nodeTest{
			data1: []interface{}{},
			data2: []interface{}{},
		})
	table = append(table,
		nodeTest{
			data1: solutions.IntToInterface(random.IntSliceBetween(0, 100)),
			data2: []interface{}{},
		})
	// ints
	for i := 0; i < 3; i++ {
		val := nodeTest{
			data1: solutions.IntToInterface(random.IntSliceBetween(0, 100)),
			data2: solutions.IntToInterface(random.IntSliceBetween(0, 100)),
		}
		table = append(table, val)
	}
	// strings
	for i := 0; i < 2; i++ {
		val := nodeTest{
			data1: solutions.StringToInterface(random.StrSlice(chars.Words)),
			data2: solutions.StringToInterface(random.StrSlice(chars.Words)),
		}
		table = append(table, val)
	}

	for _, arg := range table {
		link1 := &solutions.List{}
		linkTest := &solutions.List{}
		link2 := &student.List{}
		link2Test := &student.List{}

		for _, item := range arg.data1 {
			solutions.ListPushBack(link1, item)
			listPushBack(link2, item)
		}
		for _, item := range arg.data2 {
			solutions.ListPushBack(linkTest, item)
			listPushBack(link2Test, item)
		}

		solutions.ListMerge(link1, linkTest)
		student.ListMerge(link2, link2Test)

		solutions.ChallengeList("ListMerge", link1, copyList(link2), []interface{}{arg.data1, arg.data2})
	}
}
